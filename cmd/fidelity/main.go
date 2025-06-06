package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/jp/fidelity/internal/config"
	"github.com/jp/fidelity/internal/domain"
	fhandler "github.com/jp/fidelity/internal/handler"
	"github.com/jp/fidelity/internal/pkg/infraestructure/platform"
	"github.com/jp/fidelity/internal/repository"
)

var serverConfig config.ServerConfig

func main() {
	const ERR = 2
	var configPath string

	logger := slog.New(slog.Default().Handler())
	logger.Info("Read config")

	flag.StringVar(&configPath, "configpath", "", "File path for server configuration")

	flag.Parse()

	err := cleanenv.ReadConfig(configPath, &serverConfig)
	if err != nil {
		os.Exit(ERR)
	}

	logger.Info("Init server")
	var server *http.Server
	if serverConfig.Platform == platform.DevEnv {
		server = devEnvInject(serverConfig)
	} else {
		server = localEnvInject(serverConfig)
	}

	addShutdownHook(func(s os.Signal) {
		logger.Info("Received Signal, stopping gin server")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown: " + err.Error())
		}

		logger.Info("Shutdown has completed")
	})

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(fmt.Sprintf("Error occurred during ListenAndServe %s", err.Error()))
	}

}

func addShutdownHook(f func(s os.Signal)) {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigChannel
		f(s)
	}()
}

func common(
	cfg config.ServerConfig,
	repo domain.Repository,
	logger *slog.Logger,
) *http.Server {
	service := domain.ProviderService(repo)
	handler := fhandler.ProvideHandler(service, logger)
	server := fhandler.ProvideServer(handler, logger, cfg.Server.Host, cfg.Server.Port)

	return server
}

func devEnvInject(cfg config.ServerConfig) *http.Server {
	logger := slog.New(slog.Default().Handler())
	logger.Debug("Inject Dev Env")
	platform := platform.ProvideDevEnvPlatform(logger)
	postgresRepository := repository.ProvideGormRepository(logger, cfg.Postgres, platform)

	return common(cfg, postgresRepository, logger)
}

func localEnvInject(cfg config.ServerConfig) *http.Server {
	logger := slog.New(slog.Default().Handler())
	logger.Debug("Inject Local Env")
	platform := platform.ProvideLocalEnvPlatform(logger)
	postgresRepository := repository.ProvideGormRepository(logger, cfg.Postgres, platform)

	return common(cfg, postgresRepository, logger)
}

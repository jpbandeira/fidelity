package platform

import (
	"log/slog"
	"sync"
)

type LocalEnvPlatform struct {
	Platform
}

var (
	localEnvPlatformSvc     *LocalEnvPlatform
	localEnvPlatformSvcOnce sync.Once
)

func ProvideLocalEnvPlatform(logger *slog.Logger) *LocalEnvPlatform {
	localEnvPlatformSvcOnce.Do(func() {
		localEnvPlatformSvc = &LocalEnvPlatform{Platform: Platform{logger, LocalEnv}}
	})

	return localEnvPlatformSvc
}

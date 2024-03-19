package platform

import (
	"log/slog"
	"sync"
)

type DevEnvPlatform struct {
	Platform
}

var (
	devEnvPlatformSvc     *DevEnvPlatform
	devEnvPlatformSvcOnce sync.Once
)

func ProvideDevEnvPlatform(logger *slog.Logger) *DevEnvPlatform {
	devEnvPlatformSvcOnce.Do(func() {
		devEnvPlatformSvc = &DevEnvPlatform{Platform: Platform{logger, DevEnv}}
	})

	return devEnvPlatformSvc
}

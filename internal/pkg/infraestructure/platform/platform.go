package platform

import (
	"fmt"
	"log/slog"
)

type Platformer interface {
	GetPlatformType() Type
}

type Type string

const (
	UnknownPlatform Type = ""
	DevEnv          Type = "dev-env"
	LocalEnv        Type = "local-env"
	AWS             Type = "aws"
	COP             Type = "cop"
)

func (t Type) String() string {
	return fmt.Sprintf("<%s>", string(t))
}

type Platform struct {
	logger       *slog.Logger
	platformType Type
}

func (p *Platform) GetPlatformType() Type {
	return p.platformType
}

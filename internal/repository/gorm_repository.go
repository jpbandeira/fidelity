package repository

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jp/fidelity/internal/config"
	"github.com/jp/fidelity/internal/pkg/infraestructure/database"
	"github.com/jp/fidelity/internal/pkg/infraestructure/platform"
	"github.com/jp/fidelity/internal/repository/model"
)

// GormRepository contains the reference to the gorm DB definition and configuration
type GormRepository struct {
	*gorm.DB
	logger   *slog.Logger
	platform platform.Platformer
	cfg      config.Postgres
}

var (
	gormRepo     *GormRepository
	gormRepoOnce sync.Once
)

func databaseModels() []any {
	return []any{
		model.ServiceSummary{},
		model.Service{},
		model.Appointment{},
		model.Client{},
	}
}

// ProvideGormRepository creates the GormRepository instance
func ProvideGormRepository(logger *slog.Logger, dbConfig config.Postgres, plat platform.Platformer) *GormRepository {
	logger.Info(fmt.Sprintf("platform: %v", plat.GetPlatformType()))

	gormRepoOnce.Do(func() {
		gormRepo = newGormRepository(dbConfig, logger, plat)
	})

	return gormRepo
}

// newGormRepository Creates a new GORM repository for the specified configuration
func newGormRepository(dbConfig config.Postgres, logger *slog.Logger, plat platform.Platformer) *GormRepository {
	db, err := database.ProvideGORMDatabase(postgres.Open(dbConfig.DSN()), databaseModels())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(-1)
	}

	logger.Info("Gorm Postgres Database connection established")
	return &GormRepository{db, logger, plat, dbConfig}
}

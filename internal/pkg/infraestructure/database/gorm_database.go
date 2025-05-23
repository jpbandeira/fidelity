package database

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var errNoModelDefined = errors.New("no model defined")

func defaultGormLogger() gormlogger.Interface {
	return gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormlogger.Config{
			SlowThreshold:             1 * time.Second,
			LogLevel:                  gormlogger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
}

// ProvideGORMDatabase provides am instance of *gorm.DB that contains database connection.
func ProvideGORMDatabase(dbDialect gorm.Dialector, model []interface{}) (*gorm.DB, error) {
	if len(model) == 0 {
		return nil, errNoModelDefined
	}

	gormLogger := defaultGormLogger()
	gormConfig := gorm.Config{Logger: gormLogger}

	db, err := gorm.Open(dbDialect, &gormConfig)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(model...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

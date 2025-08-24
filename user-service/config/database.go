package config

import (
	"fmt"
	"user-service/database/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rs/zerolog/log" // sesuaikan dengan path package seeds
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dBConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PsqlDB.User,
		cfg.PsqlDB.Password,
		cfg.PsqlDB.Host,
		cfg.PsqlDB.Port,
		cfg.PsqlDB.DBName)

	db, err := gorm.Open(postgres.Open(dBConnString), &gorm.Config{})

	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to Postgres" + cfg.PsqlDB.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get database instance")
		return nil, err
	}
	seeds.SeedRole(db)
	seeds.SeedAdmin(db)

	sqlDB.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdle)

	return &Postgres{DB: db}, nil
}

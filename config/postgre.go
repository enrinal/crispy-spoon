package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Postgres struct{}

func (p *Postgres) Connect() (*gorm.DB, error) {
	pql := os.Getenv("DATABASE_URL")
	if pql == "" {
		pql = "postgresql://user:admin@localhost:54320/rg_comp_gorm?sslmode=disable"
	}

	//  DATABASE_URL=postgres://flight_api:bfAdLu2IdmEfSG5@flight-api-db.flycast:5432/flight_api?sslmode=disable

	dbConn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        pql,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func NewDB() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Reset(db *gorm.DB, table string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE " + table).Error; err != nil {
			return err
		}

		if err := tx.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1").Error; err != nil {
			return err
		}

		return nil
	})
}

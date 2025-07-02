package db

import (
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/m1sty32/sdo/pkg/config"
)

func NewPostgres(cfg *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	var db *sqlx.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			logrus.Info("Successfully connected to PostgreSQL")
			return db, nil
		}
		logrus.Warnf("Failed to connect to PostgreSQL (attempt %d/%d): %v", i+1, 10, err)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to PostgreSQL after retries: %v", err)
}
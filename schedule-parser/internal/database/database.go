package database

import (
    "fmt"
    "schedule-parser/internal/config"
    "schedule-parser/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    return db, nil
}

// RunMigrations выполняет миграции для создания таблиц
func RunMigrations(db *gorm.DB) error {
    err := db.AutoMigrate(
        &models.Teacher{},
        &models.Subject{},
        &models.Group{},
        &models.TeacherSubjectGroup{},
    )
    if err != nil {
        return fmt.Errorf("failed to run migrations: %v", err)
    }
    return nil
}
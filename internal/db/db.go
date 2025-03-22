package db

import (
    "ecommerce-app/internal/models"
    "ecommerce-app/pkg/logger"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Connect initializes a connection to the PostgreSQL database
func Connect(dsn string) (*gorm.DB, error) {
    log := logger.New()

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Error("Failed to connect to database: " + err.Error())
        return nil, err
    }

    log.Info("Successfully connected to the database")
    return db, nil
}

// Migrate runs auto-migration for the defined models
func Migrate(db *gorm.DB) error {
    log := logger.New()

    err := db.AutoMigrate(
        &models.User{},
        &models.Product{},
        &models.Cart{},
        &models.Order{},
    )
    if err != nil {
        log.Error("Failed to migrate database: " + err.Error())
        return err
    }
    
    // Create indexes for better query performance
    // Index for product name searches
    db.Exec("CREATE INDEX IF NOT EXISTS idx_products_name ON products(name)")
    // Index for product category filtering
    db.Exec("CREATE INDEX IF NOT EXISTS idx_products_category ON products(category)")
    // Index for price range queries
    db.Exec("CREATE INDEX IF NOT EXISTS idx_products_price ON products(price)")

    log.Info("Database migration and indexes creation completed")
    return nil
}

// TestConnection verifies the database connection and migrations
func TestConnection(dsn string) error {
    log := logger.New()

    // Test database connection
    db, err := Connect(dsn)
    if err != nil {
        return err
    }

    // Test migrations
    err = Migrate(db)
    if err != nil {
        return err
    }

    // Verify tables exist
    tables := []string{"users", "products", "carts", "orders"}
    for _, table := range tables {
        var count int64
        if err := db.Table(table).Count(&count).Error; err != nil {
            log.Error("Failed to verify table " + table + ": " + err.Error())
            return err
        }
        log.Info("Table " + table + " verified successfully")
    }

    log.Info("Database connection and migrations tested successfully")
    return nil
}
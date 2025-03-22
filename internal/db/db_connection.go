package db

import (
	"sync"

	"gorm.io/gorm"
)

// Global database connection instance
var (
	db     *gorm.DB
	dbOnce sync.Once
)

// SetDB sets the global database connection
func SetDB(dbConn *gorm.DB) {
	dbOnce.Do(func() {
		db = dbConn
	})
}

// GetDB returns the global database connection
func GetDB() *gorm.DB {
	return db
}
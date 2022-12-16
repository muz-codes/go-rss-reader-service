package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

// github.com/mattn/go-sqlite3
func OpenDbConnection() (*gorm.DB, error) {
	logger := zap.L()
	dbConnection, err := gorm.Open(sqlite.Open("rss_reader.sqlite"), &gorm.Config{})
	if err != nil {
		logger.Error("error in OpenDbConnection", zap.Error(err))
		return nil, err
	}
	return dbConnection, err
}

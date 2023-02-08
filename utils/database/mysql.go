package database

import (
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DefaultMysqlConnectionFromDsn returns a default GORM MySQL database
// connection from the given DSN.
func DefaultMysqlConnectionFromDsn(dsn string) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)
	dbLogMode, _ := strconv.Atoi(os.Getenv("DB_LOG_MODE"))
	if dbLogMode <= 0 {
		dbLogMode = 2
	}

	opts := &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(dbLogMode)),
	}
	return gorm.Open(dialector, opts)
}

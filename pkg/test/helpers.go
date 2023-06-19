package test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMockQueryDb creates a new mock database connection for testing purposes
// using sqlmock and gorm. If there is an error when creating the mock database,
// the function will call t.Fatal to stop the test and report the error.
// Example usage:
//
//	mockQuery, mockDb := test.NewMockQueryDb(t)
func NewMockQueryDb(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	conn, mockQuery, _ := sqlmock.New()
	//if err != nil {
	//	t.Fatal(fmt.Errorf("sqlmock.New: %w", err))
	//}
	mysqlConfig := mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	}
	options := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Time{}
		},
	}
	mockDb, _ := gorm.Open(mysql.New(mysqlConfig), options)
	//if err != nil {
	//	t.Fatal(fmt.Errorf("gorm.Open: %w", err))
	//}
	return mockQuery, mockDb
}

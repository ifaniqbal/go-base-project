package database

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestDefaultMysqlConnectionFromDsn(t *testing.T) {
	type args struct {
		dsn string
	}
	opts := &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(2)),
	}
	db, err := gorm.Open(mysql.Open(""), opts)
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name:    "error",
			args:    args{dsn: ""},
			want:    db,
			wantErr: err != nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DefaultMysqlConnectionFromDsn(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultMysqlConnectionFromDsn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

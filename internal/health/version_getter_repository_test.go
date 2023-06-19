package health

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ifaniqbal/go-base-project/pkg/test"
	"gorm.io/gorm"
)

func Test_versionGetterRepository_GetVersion(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type testCase struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDb(t)

	name := "error"
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT VERSION()")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    "",
		wantErr: true,
	})

	name = "success"
	version := "8"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"version()"}).AddRow(version),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    version,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := versionGetterRepository{
				db: tt.fields.db,
			}
			got, err := r.GetVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

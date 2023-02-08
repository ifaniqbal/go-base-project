package health

import (
	"fmt"

	"gorm.io/gorm"
)

type versionGetterRepository struct {
	db *gorm.DB
}

type VersionGetterRepository interface {
	GetVersion() (string, error)
}

// GetVersion returns the version of the database.
// It returns an error if there was a problem executing the SQL query.
func (r versionGetterRepository) GetVersion() (string, error) {
	var version string
	err := r.db.Raw("SELECT VERSION()").Scan(&version).Error
	if err != nil {
		err = fmt.Errorf("db.Scan: %w", err)
	}
	return version, err
}

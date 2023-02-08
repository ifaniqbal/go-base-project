package api

import (
	"fmt"

	"bitbucket.org/ifan-moladin/base-project/modules/health"
	"bitbucket.org/ifan-moladin/base-project/utils/catcher"
	"bitbucket.org/ifan-moladin/base-project/utils/database"
	"bitbucket.org/ifan-moladin/base-project/utils/environment"
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func Default() *Api {
	env := environment.Default()
	defaultCatcher, err := initCatcher()
	if err != nil {
		panic(fmt.Errorf("initCatcher: %w", err))
	}
	dbConn, err := initDbConn(env.Get("MYSQL_DSN"))
	if err != nil {
		panic(fmt.Errorf("initDbConn: %w", err))
	}
	routers := []Router{
		health.NewRouter(dbConn, defaultCatcher),
	}
	return New(
		httpserver.Default(),
		routers,
	)
}

type initCatcherFunc func() (catcher.Catcher, error)

func defaultInitCatcher() initCatcherFunc {
	return func() (catcher.Catcher, error) {
		defaultCatcher := catcher.Default()

		return defaultCatcher, defaultCatcher.Init()
	}
}

var initCatcher = defaultInitCatcher()

type initDbConnFunc func(string) (*gorm.DB, error)

func defaultInitDbConn() initDbConnFunc {
	return func(dsn string) (*gorm.DB, error) {
		return database.DefaultMysqlConnectionFromDsn(dsn)
	}
}

var initDbConn = defaultInitDbConn()

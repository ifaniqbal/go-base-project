package health

import (
	"bitbucket.org/ifan-moladin/base-project/utils/catcher"
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
	"gorm.io/gorm"
)

type Router struct {
	requestHandler *RequestHandler
}

func NewRouter(db *gorm.DB, defaultCatcher catcher.Catcher) *Router {
	return &Router{
		requestHandler: &RequestHandler{
			catcher: defaultCatcher,
			repo:    versionGetterRepository{db: db},
		},
	}
}

func (r Router) Route(route httpserver.RouteHandler) {
	route.GET("/", r.requestHandler.check)
}

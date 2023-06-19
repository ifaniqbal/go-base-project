package health

import (
	"github.com/ifaniqbal/go-base-project/pkg/catcher"
	"github.com/ifaniqbal/go-base-project/pkg/httpserver"
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

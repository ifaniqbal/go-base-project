package health

import (
	"github.com/ifaniqbal/go-base-project/internal/utils"
	"gorm.io/gorm"
)

type Router struct {
	requestHandler *RequestHandler
}

func NewRouter(db *gorm.DB, defaultCatcher utils.Catcher) *Router {
	return &Router{
		requestHandler: &RequestHandler{
			catcher: defaultCatcher,
			repo:    versionGetterRepository{db: db},
		},
	}
}

func (r Router) Route(route utils.RouteHandler) {
	route.GET("/", r.requestHandler.check)
}

package api

import "github.com/ifaniqbal/go-base-project/internal/utils"

type Api struct {
	httpServer utils.HttpServer
	routers    []Router
}

type Router interface {
	Route(rgh utils.RouteHandler)
}

func (a Api) Start() error {
	root := a.httpServer.Group("/")
	for _, router := range a.routers {
		router.Route(root)
	}

	var err error
	if err = a.httpServer.Run(); err != nil {
		return err
	}

	return nil
}

func New(
	httpServer utils.HttpServer,
	routers []Router,
) *Api {
	return &Api{
		httpServer: httpServer,
		routers:    routers,
	}
}

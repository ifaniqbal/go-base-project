// Package httpserver provides a simple HTTP server implementation
// using Gin framework.
package httpserver

import (
	"github.com/gin-gonic/gin"
)

// Gin type is a struct that contains a pointer to a Gin engine.
type Gin struct {
	engine *gin.Engine
}

// New creates a new instance of the Gin http server with the provided gin.Engine.
func New(engine *gin.Engine) *Gin {
	return &Gin{
		engine: engine,
	}
}

// Default function creates a Gin engine with the default middleware
// and returns a pointer to a Gin type.
func Default() *Gin {
	engine := gin.Default()
	return New(engine)
}

// Run runs the Gin engine with the provided addresses.
func (g Gin) Run(addr ...string) error {
	return g.engine.Run(addr...)
}

// Group creates a new router group with the provided relativePath
// and handles as middlewares. It returns a RouteHandler type.
func (g Gin) Group(relativePath string, handles ...HandlerFunc) RouteHandler {
	routerGroup := g.engine.Group(relativePath, toGinHandlerFuncs(handles)...)
	return &GinRouteHandler{routerGroup: routerGroup}
}

// toGinHandlerFuncs converts []HandlerFunc to []gin.HandlerFunc.
func toGinHandlerFuncs(handles []HandlerFunc) []gin.HandlerFunc {
	ginHandles := make([]gin.HandlerFunc, len(handles))
	for i, handle := range handles {
		ginHandles[i] = toGinHandlerFunc(handle)
	}
	return ginHandles
}

// toGinHandlerFunc takes in a HandlerFunc and returns a gin.HandlerFunc.
func toGinHandlerFunc(handle HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handle(ctx)
	}
}

// GinRouteHandler represents a Gin Router Group, and
// implements the RouteHandler interface.
type GinRouteHandler struct {
	routerGroup *gin.RouterGroup
}

// Group creates a new child router group with the provided relativePath and
// handles as middlewares. It returns a RouteHandler type.
func (grh GinRouteHandler) Group(relativePath string, handles ...HandlerFunc) RouteHandler {
	routerGroup := grh.routerGroup.Group(relativePath, toGinHandlerFuncs(handles)...)
	return &GinRouteHandler{routerGroup: routerGroup}
}

// BasePath returns the base path of the Gin Router Group.
func (grh GinRouteHandler) BasePath() string {
	return grh.routerGroup.BasePath()
}

// Use adds the provided list of HandlerFunc as middlewares to the Gin Router Group.
// It returns the RouteHandler type.
func (grh GinRouteHandler) Use(middlewares ...HandlerFunc) RouteHandler {
	grh.routerGroup.Use(toGinHandlerFuncs(middlewares)...)
	return grh
}

// Handle creates a new route with the provided httpMethod, relativePath, and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) Handle(
	httpMethod string,
	relativePath string,
	handles ...HandlerFunc,
) RouteHandler {
	grh.routerGroup.Handle(httpMethod, relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

// GET creates a new GET route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) GET(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.GET(relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

// POST creates a new POST route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) POST(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.POST(relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

// DELETE creates a new DELETE route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) DELETE(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.DELETE(relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

// PATCH creates a new PATCH route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) PATCH(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.PATCH(relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

// PUT creates a new PUT route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) PUT(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.PUT(relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

// OPTIONS creates a new OPTIONS route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) OPTIONS(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.OPTIONS(relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

// HEAD creates a new HEAD route with the provided relativePath and
// handles. It returns the RouteHandler type.
func (grh GinRouteHandler) HEAD(relativePath string, handles ...HandlerFunc) RouteHandler {
	grh.routerGroup.HEAD(relativePath, toGinHandlerFuncs(handles)...)
	return grh
}

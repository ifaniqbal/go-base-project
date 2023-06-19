package health

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ifaniqbal/go-base-project/pkg/catcher"
	"github.com/ifaniqbal/go-base-project/pkg/httpserver"
)

type RequestHandler struct {
	repo    VersionGetterRepository
	catcher catcher.Catcher
}

func (h RequestHandler) check(ctx httpserver.Context) {
	currentCatcher := h.catcher.WithContext(ctx)
	version, err := h.repo.GetVersion()
	if err != nil {
		err = fmt.Errorf("repo.GetVersion: %w", err)
		currentCatcher.Catch(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "System is healthy",
		"mysqlVersion": version,
	})
}

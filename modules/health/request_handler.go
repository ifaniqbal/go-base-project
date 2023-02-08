package health

import (
	"fmt"
	"net/http"

	"bitbucket.org/ifan-moladin/base-project/utils/catcher"
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
	"github.com/gin-gonic/gin"
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

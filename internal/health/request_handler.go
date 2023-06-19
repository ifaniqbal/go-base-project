package health

import (
	"fmt"
	"github.com/ifaniqbal/go-base-project/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	repo    VersionGetterRepository
	catcher utils.Catcher
}

func (h RequestHandler) check(ctx utils.Context) {
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

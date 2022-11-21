package Home

import (
	"net/http"

	"github.com/EdwBaeza/inhouse/internal"
	"github.com/gin-gonic/gin"
)

// ShowHandler a single home
func ShowHandler(repository internal.HomeRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		home := repository.Find("test")
		ctx.JSON(http.StatusOK, FromHomeToReq(home))
	}
}

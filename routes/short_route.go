package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/services"
)

func CreateShortRoute(s services.ShortService, e *gin.Engine) {
	e.GET("/:tag", func(ctx *gin.Context) {
		tag := ctx.Param("tag")
		url, err := s.GetURL(tag)
		if err != nil {
			ctx.String(http.StatusNotFound, "URL not found")
			return
		}
		ctx.Redirect(http.StatusFound, *url)
	})
}
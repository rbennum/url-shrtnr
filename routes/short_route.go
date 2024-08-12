package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/services"
)

func CreateShortRoute(s services.ShortService, e *gin.Engine) {
	// show main html
	e.GET("/", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"main.html",
			gin.H {
				"status": "success",
			},
		)
	})
}
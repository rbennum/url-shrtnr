package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/models"
	"github.com/rbennum/url-shrtnr/services"
)

func CreateMainRoute(s services.ShortService, e *gin.Engine) {
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
	// create a new shorter version of a URL
	e.POST("/url", func(ctx *gin.Context) {
		var req models.LinkRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H {"error": err.Error()},
			)
			return
		}
		log.Printf("[Body] %s: %v", ctx.Request.URL.String(), req)
		url_obj, err := s.CreateURL(req.URL)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest, 
				gin.H {"error": err.Error()},
			)
			return
		}
		ctx.JSON(http.StatusCreated, url_obj)
	})
}
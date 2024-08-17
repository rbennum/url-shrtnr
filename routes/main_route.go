package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/models"
	"github.com/rbennum/url-shrtnr/services"
)

type MainRoute struct {
	service services.ShortService
}

func NewMainRoute(s *services.ShortService) *MainRoute {
	return &MainRoute{
		service: *s,
	}
}

func CreateMainRoute(m *MainRoute, e *gin.Engine) {
	e.GET("/", m.OpenMainPage)
	e.POST("/url", m.CreateShortURL)
}

// show main html
func (r *MainRoute) OpenMainPage(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"main.html",
		gin.H{
			"status": "success",
		},
	)
}

// create a new shorter version of a URL
func (r *MainRoute) CreateShortURL(ctx *gin.Context) {
	var req models.LinkRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	log.Printf("[Body] %s: %v", ctx.Request.URL.String(), req)
	url_obj, err := r.service.CreateURL(req.URL)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	ctx.JSON(http.StatusCreated, url_obj)
}

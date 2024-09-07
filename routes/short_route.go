package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/services"
)

type ShortRouter struct {
	service services.ShortService
}

func NewShortRouter(s *services.ShortService) *ShortRouter {
	return &ShortRouter{
		service: *s,
	}
}

func CreateShortRoute(r *ShortRouter, e *gin.Engine) {
	e.GET("/", r.RedirectURL)
	e.GET("/:tag", r.RedirectURL)
}

func (r *ShortRouter) RedirectURL(ctx *gin.Context) {
	tag := ctx.Param("tag")
	if tag == "" {
		log.Println("No tag available")
		r.redirectToNotFound(ctx)
		return
	}
	url, err := r.service.GetURL(tag)
	if err != nil {
		log.Printf("[Tag %s] not found", tag)
		r.redirectToNotFound(ctx)
		return
	}
	if !strings.HasPrefix(*url, "https://") && !strings.HasPrefix(*url, "http://") {
		*url = "https://" + *url
	}
	ctx.Redirect(http.StatusFound, *url)
}

func (r *ShortRouter) redirectToNotFound(ctx *gin.Context) {
	ctx.HTML(
		http.StatusNotFound,
		"not-found.html",
		gin.H{
			"status": "not found",
		},
	)
}

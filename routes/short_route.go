package routes

import (
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
	e.GET("/:tag", r.RedirectURL)
}

func (r *ShortRouter) RedirectURL(ctx *gin.Context) {
	tag := ctx.Param("tag")
	url, err := r.service.GetURL(tag)
	if err != nil {
		ctx.String(http.StatusNotFound, "URL not found")
		return
	}
	if !strings.HasPrefix(*url, "https://") && !strings.HasPrefix(*url, "http://") {
		*url = "https://" + *url
	}
	ctx.Redirect(http.StatusFound, *url)
}

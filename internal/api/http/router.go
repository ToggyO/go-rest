package api

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type AppRouter struct {
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

func NewAppRouter() *AppRouter {
	r := gin.Default()
	return &AppRouter{
		Router:      r,
		RouterGroup: &r.RouterGroup,
	}
}

func (e *AppRouter) AddGlobalRoutePrefix(prefix string) {
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	e.RouterGroup = e.Router.Group(prefix)
}

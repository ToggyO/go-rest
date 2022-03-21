package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Build() http.Handler {
	router := gin.Default()
	return router
}

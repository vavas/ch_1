package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vavas/challenge/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/email", h.addEmail)
	router.GET("/emails", h.getAllEmails)
	router.DELETE("/email/:email", h.removeEmail)

	return router
}

package delivery

import (
	"github.com/gin-gonic/gin"
	"vk-task/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		adverts := api.Group("/adverts")
		{
			adverts.POST("/", h.createAdvert)
			adverts.GET("/", h.listAdverts)
		}
	}
	return router
}

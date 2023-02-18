package handler

import (
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/usecase/notify_group"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase notify_group.Usecase
}

func NewHandler(uc notify_group.Usecase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/sign-up", h.signUp)
	// 	auth.POST("/sign-in", h.signIn)
	// }

	api := router.Group("/api")
	{
		notifyGroups := api.Group("notifyGroups")
		{
			notifyGroups.POST("/", h.CreateNotifyGroup)
		}
	}

	return router
}

func (h *Handler) SetRoutes(g *gin.RouterGroup) {
	notifyGroup := g.Group("/NotifyGroup")

	notifyGroup.POST("/", h.CreateNotifyGroup)

}

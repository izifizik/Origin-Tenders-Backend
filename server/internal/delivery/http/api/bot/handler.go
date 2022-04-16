package bot

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"origin-tender-backend/server/internal/delivery/http/api"
	botService "origin-tender-backend/server/internal/service/bot-service"
)

type handler struct {
	service botService.BotService
}

func NewBotHandler(service botService.BotService) api.Handler {
	return &handler{service}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST("/bot/generateToken", h.GenerateToken)
}

func (h *handler) GenerateToken(c *gin.Context) {
	// принимает тип(ник\груп айди)
	var dto TokenDTO

	err := c.ShouldBindJSON(&dto)

	token := h.service.GenerateToken(dto.Type)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	// отдает токен
}

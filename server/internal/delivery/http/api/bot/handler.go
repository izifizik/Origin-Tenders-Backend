package bot

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"origin-tender-backend/server/internal/delivery/http/api"
	"origin-tender-backend/server/internal/domain"
	botService "origin-tender-backend/server/internal/service/bot-service"
	"time"
)

type handler struct {
	botService botService.BotService
	//tenderService tenderService.TenderService
}

func NewBotHandler(service botService.BotService) api.Handler {
	return &handler{service}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST("/bot/generateToken", h.GenerateToken)
	router.GET("/tender/:tenderID", h.GetTender)
	router.GET("/user/:uuid", h.GetUser)
	//router.POST("/bot/proofToken")
}

func (h *handler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ID":            primitive.NewObjectID(),
		"Name":          "Тетя Вася",
		"Notifications": domain.Notifications{TgID: "asd"},
		"Filters":       []string{"filter1", "filter2"},
		"TendersHistory": []domain.Tender{
			{
				ID:                 primitive.NewObjectID(),
				Name:               "Tender1",
				TimeEnd:            time.Now().Add(time.Hour * 24),
				Description:        "asdasdasdasd asd as da sd asd ",
				Filters:            []string{"tag1", "tag2"},
				StartPrice:         123123123.123,
				CurrentPrice:       1.11,
				Status:             "Open??? nujen li on",
				MinimalStepPercent: 0.5,
				MaxStepPercent:     1.1,
			}, {
				ID:                 primitive.NewObjectID(),
				Name:               "Tender2",
				TimeEnd:            time.Now().Add(time.Hour * 24),
				Description:        "asdasdasdasd asd as da sd asd ",
				Filters:            []string{"tag1", "tag2"},
				StartPrice:         123123123.123,
				CurrentPrice:       1.11,
				Status:             "Open??? nujen li on",
				MinimalStepPercent: 0.5,
				MaxStepPercent:     1.1,
			},
		},
	})
}

func (h *handler) GetTender(c *gin.Context) {
	// tenderh.tenderService.GetByID(c.Param("tenderID"))
	tenders := []domain.Tender{
		{
			ID:                 primitive.NewObjectID(),
			Name:               "Tender1",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender2",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender3",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender4",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender5",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender6",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender7",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender8",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender9",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender10",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender112",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender12",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender13",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender14",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender115",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender16",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender17",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender18",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender19",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		}, {
			ID:                 primitive.NewObjectID(),
			Name:               "Tender20",
			TimeEnd:            time.Now().Add(time.Hour * 24),
			Description:        "asdasdasdasd asd as da sd asd ",
			Filters:            []string{"tag1", "tag2"},
			StartPrice:         123123123.123,
			CurrentPrice:       1.11,
			Status:             "Open??? nujen li on",
			MinimalStepPercent: 0.5,
			MaxStepPercent:     1.1,
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"tenders": tenders,
		"page":    1,
	})
}

func (h *handler) GenerateToken(c *gin.Context) {
	// принимает тип(ник\груп айди)
	var dto TokenDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//token := h.botService.GenerateToken(dto.Type)

	c.JSON(http.StatusOK, gin.H{
		"token": "jqheui321ye98",
	})
	// отдает токен
}

//func (h *handler) ProofToken(c *gin.Context) {
//	var dto TokenProofDTO
//
//	err := c.ShouldBindJSON(&dto)
//	if err != nil {
//		c.AbortWithStatus(http.StatusBadRequest)
//		return
//	}
//
//	//err := h.botService.ProofToken(dto.ID, dto.Token)
//
//	c.JSON()
//}

package bot

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"net/http"
	"origin-tender-backend/server/internal/delivery/http/api"
	"origin-tender-backend/server/internal/domain"
	botService "origin-tender-backend/server/internal/service/bot-service"
	"strconv"
	"time"
)

type handler struct {
	botService botService.BotService
	// tenderService tenderService.TenderService
	// notificationService
}

func NewBotHandler(service botService.BotService) api.Handler {
	return &handler{service}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST("/bot/generateToken", h.GenerateToken)
	router.GET("/tenders", h.GetTenders)
	router.GET("/user/:uuid", h.GetUser)
	router.POST("/bot/set_options")
	router.GET("/ws/notifications", h.Notification)
	//router.POST("/bot/proofToken")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *handler) Notification(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// читать из канала range и записывать в вебсокет (сейчас это фекалия)
	// Кеша же гений веб сокетов пускай делает
	go func() { // в нотификейшоне надо что бы мы отправляли увед и в коннект (если тот открыт) и в бота (всегда)
		conn.WriteJSON(gin.H{
			"id":            "3",
			"type":          "participation",
			"description":   "Бот сделал ставку в тендере ХХХ",
			"tenderID":      "3",
			"price":         2212,
			"isNeedApprove": true,
			"isApproved":    true,
			"approve":       false,
		})
	}()
	// отправка в бота (хз как)
}

func (h *handler) BotSetOptions(c *gin.Context) {
	// создаем горутину которая крутит этого бота в тендере с определенными настройками
	var dto OptionsDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	go BotStart(dto)
}

func (h *handler) Approve(c *gin.Context) {
	// когда автоматически бот уже не может подтверждать сделки будет вызыватсья этот поинт
	// Работает следующим образом
	// Бот в фоне ждет подтверждения => надо реализовать где то подтверждение транзакции

}

func (h *handler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ID":      primitive.NewObjectID(),
		"Name":    "Тетя Вася",
		"Filters": []string{"filter1", "filter2"},
		"TendersHistory": []domain.Tender{
			{
				ID:           primitive.NewObjectID(),
				Name:         "Tender1",
				TimeEnd:      time.Now().Add(time.Hour * 24),
				Description:  "asdasdasdasd asd as da sd asd ",
				Filters:      []string{"tag1", "tag2"},
				StartPrice:   123123123.123,
				CurrentPrice: 1.11,
				Status:       "Open??? nujen li on",
				StepPercent:  0.5,
			}, {
				ID:           primitive.NewObjectID(),
				Name:         "Tender2",
				TimeEnd:      time.Now().Add(time.Hour * 24),
				Description:  "asdasdasdasd asd as da sd asd ",
				Filters:      []string{"tag1", "tag2"},
				StartPrice:   123123123.123,
				CurrentPrice: 1.11,
				Status:       "Open??? nujen li on",
				StepPercent:  0.5,
			},
		},
	})
}

func (h *handler) GetTenders(c *gin.Context) {
	// tenderh.tenderService.GetByID(c.Param("tenderID"))
	tenders := []domain.Tender{
		{
			ID:           primitive.NewObjectID(),
			Name:         "Tender1",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender2",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender3",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender4",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender5",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender6",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender7",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender8",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender9",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender10",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender112",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender12",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender13",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender14",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender115",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender16",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender17",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender18",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender19",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
		}, {
			ID:           primitive.NewObjectID(),
			Name:         "Tender20",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "asdasdasdasd asd as da sd asd ",
			Filters:      []string{"tag1", "tag2"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open??? nujen li on",
			StepPercent:  0.5,
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
	rand.Seed(time.Now().Unix())
	c.JSON(http.StatusOK, gin.H{
		"token": strconv.Itoa(rand.Int()),
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

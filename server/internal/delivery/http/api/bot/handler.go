package bot

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"net/http"
	"origin-tender-backend/server/internal/delivery/http/api"
	"origin-tender-backend/server/internal/domain"
	botService "origin-tender-backend/server/internal/service/bot-service"
	"origin-tender-backend/server/internal/service/teleg-bot-service/actions"
	"origin-tender-backend/websocket/wsActions"
	"strconv"
	"time"
)

type handler struct {
	botService botService.BotService
	//tenderService tenderService.TenderService
}

func NewBotHandler(service botService.BotService) api.Handler {
	return &handler{service}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.Use(CORSMiddleware())
	router.POST("/bot/generateToken", func(context *gin.Context) {
		h.GenerateToken(context)
	})
	router.POST("/bot/generateToken2", func(context *gin.Context) {
		h.GenerateToken2(context, h.botService)
	})
	router.GET("/tenders", h.GetTenders)
	router.GET("/tender/:tender_id", h.GetTender)
	router.GET("/user/:uuid", h.GetUser)
	router.POST("/bot/set_options")
	router.GET("/ws/notification", h.NotificationWS)
	//router.POST("/bot/proofToken")

	router.POST("/order", func(context *gin.Context) {
		h.CreateOrder(context, h.botService)
	})

	router.POST("/test/event", func(context *gin.Context) {
		h.RaiseEvent(context, h.botService)
	})

}

func (h *handler) SetupBot(c *gin.Context) {
	var dto BotSetupDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	h.botService.BotSetup(dto.UserID, dto.TenderID, dto.Alg, dto.Type, dto.Procent, dto.Minimal, dto.Critical, dto.IsApprove)

	c.Status(http.StatusOK)
}

func (h *handler) RaiseEvent(c *gin.Context, s botService.BotService) {
	var event domain.ServiceEvent
	c.ShouldBindJSON(&event)

	if event.Type == "tender" {
		var tender domain.Tender
		err := json.Unmarshal([]byte(event.Data), &tender)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(tender.Name)

		users, err := s.GetTgUsers()

		for _, u := range users {
			err := actions.SendAcceptParticipationInTender(u.UserId, tender)
			if err != nil {
				fmt.Println(err)
			}
		}

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(len(users))
		}

		s.CreateTender(tender)

		wsActions.NotifyAllSession("sess")

	} else if event.Type == "order" {
		var order domain.Order
		err := json.Unmarshal([]byte(event.Data), &order)
		if err != nil {
			fmt.Println(err)
		}

		err = s.CreateOrder(order)
		if err != nil {
			fmt.Println(err)
		}

		orderData, err := json.Marshal(order)
		if err != nil {
			fmt.Println(err)
		}

		//TODO: do stuff

		wsActions.NotifyUser(string(orderData), order.UserId)
		wsActions.NotifyAllBet(string(orderData))

	}

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *handler) NotificationWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// go + уведы в тг бота
	h.botService.SentNotification(conn)
}

func (h *handler) CreateOrder(c *gin.Context, s botService.BotService) {
	var order domain.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		fmt.Println(err)
	}

	err = s.CreateOrder(order)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{})
}

// какую статистику стоит показывать по тендерам в которых учавствует бот от /tenders
// это активные тендеры бота

func (h *handler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ID":      primitive.NewObjectID(),
		"Name":    "Тетя Вася",
		"Filters": []string{"filter1", "filter2"},
		"TendersHistory": []domain.Tender{
			{
				ID:           "1",
				Name:         "Tender1",
				TimeEnd:      time.Now().Add(time.Hour * 24),
				Description:  "asdasdasdasd asd as da sd asd ",
				Filters:      []string{"tag1", "tag2"},
				StartPrice:   123123123.123,
				CurrentPrice: 1.11,
				Status:       "Open??? nujen li on",
				StepPercent:  0.5,
			}, {
				ID:           "2",
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

func (h *handler) GetTender(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ID":               c.Param("tender_id"),
		"Name":             "Закуп пиломатериала из древесины",
		"TimeEnd":          time.Now().Add(time.Hour * 24),
		"Description":      "Древесина хвойных пород не ниже 3 сорта по ГОСТ 8486 и не ниже 2 сорта лиственных пород по ГОСТ 2695",
		"ShortDescription": "Брус, доска обрезная /необрезная. Количество пиломатериала - 500 тыс.",
		"Filters":          []string{"Фильтр по цене"},
		"StartPrice":       2300050.50,
		"CurrentPrice":     2145600.45,
		"Status":           "Активна",
		"StepPercent":      0.5,
	})
}

func (h *handler) GetTenders(c *gin.Context) {
	// tenderh.tenderService.GetByID(c.Param("tenderID"))

	// дескрипцию для тендеров нормальную сделать
	// занести эти тендеры в бд и доставать оттуда и отдавать
	//tenders := h.botService.GetAllTenders()
	tenders := []domain.Tender{
		{
			ID:           "1",
			Name:         "Покупка разеток тип С",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "Металлческие листы: \n Лист 10мм г/к ГОСТ 1577-93 1500*6000 \n Лист 5мм г/к ГОСТ 14637-89 1500*6000 \n Лист 8мм г/к ГОСТ 14637-89 1500*3000",
			Filters:      []string{"Price filter"},
			StartPrice:   5000000,
			CurrentPrice: 4850000.50,
			Status:       "Активна",
			StepPercent:  0.5,
		}, {
			ID:           "2",
			Name:         "Tender Name",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "link to tender: http://10.10.117.179:8080/tender/" + "2",
			Filters:      []string{"Some filter"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open",
			StepPercent:  0.5,
		}, {
			ID:           "3",
			Name:         "Tender Name",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "link to tender: http://10.10.117.179:8080/tender/" + "3",
			Filters:      []string{"Some filter"},
			StartPrice:   123123123.123,
			CurrentPrice: 1.11,
			Status:       "Open",
			StepPercent:  0.5,
		}, {
			ID:               "4",
			Name:             "Tender Name",
			TimeEnd:          time.Now().Add(time.Hour * 24),
			Description:      "Древесина хвойных пород не ниже 3 сорта по ГОСТ 8486 и не ниже 2 сорта лиственных пород по ГОСТ 2695",
			ShortDescription: "Брус, доска обрезная /необрезная. Количество пиломатериала - 500 тыс.",
			Filters:          []string{"Price filter"},
			StartPrice:       123123123.123,
			CurrentPrice:     1.11,
			Status:           "Активна",
			StepPercent:      0.5,
		},
		{
			ID:           "5",
			Name:         "Tender Name",
			TimeEnd:      time.Now().Add(time.Hour * 24),
			Description:  "link to tender: http://10.10.117.179:8080/tender/" + "5",
			Filters:      []string{"Some filter"},
			StartPrice:   123123.123,
			CurrentPrice: 1.11,
			Status:       "Open",
			StepPercent:  1,
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"tenders": tenders,
		"page":    1,
	})
}

func (h *handler) GenerateToken2(c *gin.Context, s botService.BotService) {
	// принимает тип(ник\груп айди)
	var dto TokenDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//token := h.botService.GenerateToken(dto.Type)
	rand.Seed(time.Now().Unix())

	seed := strconv.Itoa(rand.Int())

	user, err := s.GetSiteUserByName(dto.Value)

	err = s.CreateTgToken(user.Name, seed, user.ID)

	c.JSON(http.StatusOK, gin.H{
		"token": seed,
	})
	// отдает токен
}

func (h *handler) GenerateToken(c *gin.Context) {
	// принимает тип(ник\груп айди)
	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var dto TokenDTO

	err = c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//token := h.botService.GenerateToken(dto.Type)
	rand.Seed(time.Now().Unix())
	c.JSON(http.StatusOK, gin.H{
		"token": rand.Int(),
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

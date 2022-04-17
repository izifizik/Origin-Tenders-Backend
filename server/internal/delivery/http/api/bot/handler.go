package bot

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"origin-tender-backend/server/internal/delivery/http/api"
	"origin-tender-backend/server/internal/domain"
	botService "origin-tender-backend/server/internal/service/bot-service"
	"origin-tender-backend/server/internal/service/teleg-bot-service/actions"
	"origin-tender-backend/server/internal/service/wsActions"
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
	router.POST("/bot/generateToken2", h.GenerateToken)

	router.POST("/order", h.CreateOrder)
	router.GET("/tenders", h.GetTenders)
	router.GET("/tender/:tender_id", h.GetTender)

	router.GET("/ws/bets/:id", h.Bets)
	router.GET("/ws/notify/:id", h.Notify)
	router.GET("/ws/session/:id", h.Session)

	router.POST("/test/event", h.RaiseEvent)
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

func (h *handler) RaiseEvent(c *gin.Context) {
	var event domain.ServiceEvent
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if event.Type == "tender" {
		var tender domain.Tender
		err = json.Unmarshal([]byte(event.Data), &tender)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(tender.Name)

		users, err := h.botService.GetTgUsers()
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		fmt.Println(len(users))

		for _, u := range users {
			err = actions.SendAcceptParticipationInTender(u.UserId, tender)
			if err != nil {
				fmt.Println(err)
			}
		}

		err = h.botService.CreateTender(tender)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		wsActions.NotifyAllSession("sess")
		c.Status(http.StatusOK)
		return
	} else if event.Type == "order" {

		var order domain.Order
		err = json.Unmarshal([]byte(event.Data), &order)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = h.botService.CreateOrder(order)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		orderData, err := json.Marshal(order)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		wsActions.NotifyUser(string(orderData), order.UserId)
		wsActions.NotifyAllBet(string(orderData))
	} else {
		fmt.Println("uncnown type!")
	}

}

func (h *handler) CreateOrder(c *gin.Context) {
	var order domain.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.botService.CreateOrder(order)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(200)
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

func (h *handler) GenerateToken(c *gin.Context) {
	var dto TokenDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().Unix())

	seed := strconv.Itoa(rand.Int())

	user, err := h.botService.GetSiteUserByName(dto.Value)

	err = h.botService.CreateTgToken(user.Name, seed, user.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": seed,
	})
}

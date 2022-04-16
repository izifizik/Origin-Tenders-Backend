package botService

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"math/rand"
	"origin-tender-backend/server/internal/domain"
	"origin-tender-backend/server/internal/repository/mongodb"
	"strconv"
)

type service struct {
	repo         mongodb.Repository
	notification chan domain.Notification
}

func NewBotService(repo mongodb.Repository) BotService {
	n := make(chan domain.Notification)
	return &service{repo: repo, notification: n}
}

func NewStandardBot(id, tenderID string, procent, critical float64, isApprove bool) {

}

func NewExtendedBot(id, tenderID string, procent, minimal, critical float64, isApprove bool) {

}

func (s *service) BotSetup(id, tenderID, alg, tpe string, procent, minimal, critical float64, isApprove bool) {
	switch alg {
	case "to_small":
		if tpe == "standard" {
			go NewStandardBot(id, tenderID, procent, critical, isApprove)
		}
		go NewExtendedBot(id, tenderID, procent, minimal, critical, isApprove)
	case "race_win":

	case "curr_procent":

	}
}

func (s *service) BotActivate(id, tenderId string, stepPercent, criticalPrice float64, isNeedApprove bool) {
	// бот создает участие (фиксация в бд) на определенный юзер айди с опциями
	//s.repo.CreateBotByID(id, tenderId, stepPercent, criticalPrice, isNeedApprove)
	//// бот в горутине
	//go func() {
	//	// бот получает цену из бд по тендер айди
	//	// бот смотрит не был ли он последним кто менял цену
	//	// бот меняет цену на определенный шаг
	//	// бот записывает данные в бд
	//	// бот ждет 25 секу
	//}()
}

func (s *service) SentNotification(conn *websocket.Conn) {
	for notification := range s.notification {
		data, err := json.Marshal(&notification)
		if err != nil {
			log.Println("json marshal: " + err.Error())
			continue
		}
		err = conn.WriteMessage(1, data)
		if err != nil {
			log.Println("write message ws: " + err.Error())
			continue
		}
	}
}

func (s *service) WriteNotification(notification domain.Notification) {
	s.notification <- notification
}

func (s *service) GetSiteUserByName(name string) (domain.User, error) {
	return s.repo.GetSiteUserByName(name)
}

func (s *service) CreateSiteUser(user domain.User) error {
	return s.repo.CreateSiteUser(user)
}

// siteId - User collection ObjectId
func (s *service) CreateTgToken(name string, token string, siteId primitive.ObjectID) error {
	return s.repo.CreateToken(name, token, siteId)
}

func (s *service) GenerateToken(ID string) string {
	h := md5.New()
	io.WriteString(h, ID)
	token := generateToken(binary.BigEndian.Uint64(h.Sum(nil)))
	s.repo.SaveToken(ID, token)
	return token
}

func (s *service) ProofToken(ID string, token string) error {
	return s.repo.ProofToken(ID, token)
}

func generateToken(seed uint64) string {
	rand.Seed(int64(seed))
	return strconv.Itoa(rand.Int())
}

//func (s *service) StartServeTendor()

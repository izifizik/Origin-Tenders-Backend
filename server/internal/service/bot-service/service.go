package botService

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"math/rand"
	"origin-tender-backend/server/internal/domain"
	"origin-tender-backend/server/internal/repository/mongodb"
	"strconv"
	"time"
)

type service struct {
	repo         mongodb.Repository
	notification chan domain.Notification
}

func NewBotService(repo mongodb.Repository) BotService {
	n := make(chan domain.Notification)
	return &service{repo: repo, notification: n}
}

func (s *service) BotSetup(id, tenderID, alg, tpe string, procent, minimal, critical float64, isApprove bool) {
	switch alg {
	case "to_small":
		if tpe == "standard" {
			for { // правильно конечно делать это с получением извне но и фор пока что тоже выглядит не плохо ))))))))))))))))))))))))))))))))))))))))))))))))))))))))
				tender := s.repo.GetTenderByID(tenderID)
				if tender.Owner == id {
					time.Sleep(time.Second * 4)
					continue
				}

				if tender.CurrentPrice < critical {
					break
				}

				updatePrice := (tender.CurrentPrice * tender.StepPercent) - tender.CurrentPrice

				order := domain.Order{
					TimeStamp:  time.Now(),
					UserId:     id,
					TenderId:   tenderID,
					TenderName: tender.Name,
					Price:      updatePrice,
				}

				err := s.repo.CreateOrder(order)
				if err != nil {
					fmt.Println("error with order create: " + err.Error())
					continue
				}
				time.Sleep(4 * time.Second)
			}
			break
		}

		for {
			tender := s.repo.GetTenderByID(tenderID)
			if tender.Owner == id {
				time.Sleep(4 * time.Second)
				continue
			}

			if tender.CurrentPrice < minimal {
				//sent notification and approve
				if isApprove == false {
					break
				}
			}

			if tender.CurrentPrice < critical {
				break
			}

			updatePrice := (tender.CurrentPrice * tender.StepPercent) - tender.CurrentPrice

			order := domain.Order{
				TimeStamp:  time.Now(),
				UserId:     id,
				TenderId:   tenderID,
				TenderName: tender.Name,
				Price:      updatePrice,
			}

			err := s.repo.CreateOrder(order)
			if err != nil {
				fmt.Println("error with order create: " + err.Error())
				continue
			}
			time.Sleep(4 * time.Second)
		}
	case "race_win": // подтверждение такого алго отправляется вместе с отправкой настройки и больше никогда
		if isApprove == false {
			break
		}
		for {
			tender := s.repo.GetTenderByID(tenderID)

			if tender.Owner == id {
				time.Sleep(4 * time.Minute)
				continue
			}

			updatePrice := (tender.CurrentPrice * tender.StepPercent) - tender.CurrentPrice

			order := domain.Order{
				TimeStamp:  time.Now(),
				UserId:     id,
				TenderId:   tenderID,
				TenderName: tender.Name,
				Price:      updatePrice,
			}

			err := s.repo.CreateOrder(order)
			if err != nil {
				fmt.Println("error with order create: " + err.Error())
				continue
			}
			time.Sleep(4 * time.Second)
		}
	case "curr_procent":
		for {
			tender := s.repo.GetTenderByID(tenderID)

			if tender.Owner == id {
				time.Sleep(4 * time.Second)
				continue
			}

			if procent > (1-(tender.CurrentPrice/tender.StartPrice))*100 {
				break
			}

			updatePrice := (tender.CurrentPrice * tender.StepPercent) - tender.CurrentPrice
			order := domain.Order{
				TimeStamp:  time.Now(),
				UserId:     id,
				TenderId:   tenderID,
				TenderName: tender.Name,
				Price:      updatePrice,
			}

			err := s.repo.CreateOrder(order)
			if err != nil {
				fmt.Println("error with order create: " + err.Error())
				continue
			}
			time.Sleep(4 * time.Second)
		}
	}
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

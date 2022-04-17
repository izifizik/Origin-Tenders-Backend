package botService

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"origin-tender-backend/server/internal/domain"
	"origin-tender-backend/server/internal/repository/mongodb"
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

// че то не так как зауманно
func (s *service) BotSetup(id, tenderID, alg, tpe string, procent, minimal, critical float64, isApprove bool) {
	switch alg {
	case "to_small":
		if tpe == "standard" {
			for { // правильно конечно делать это с получением извне но и фор пока что тоже выглядит не плохо ))))))))))))))))))))))))))))))))))))))))))))))))))))))))
				log.Println("user: " + id + ", alg: " + alg)

				elementOfSurprise := time.Duration(rand.Intn(5))

				tender, err := s.repo.GetTenderByID(tenderID)
				if err != nil {
					continue
				}
				if tender.Owner == id {
					time.Sleep(time.Second * (4 + elementOfSurprise))
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

				err = s.repo.CreateOrder(order)
				if err != nil {
					fmt.Println("error with order create: " + err.Error())
					continue
				}
				time.Sleep(4 * (time.Second + elementOfSurprise))
			}
			break
		}

		for {
			log.Println("user: " + id + ", alg: " + alg)

			elementOfSurprise := time.Duration(rand.Intn(5))

			tender, err := s.repo.GetTenderByID(tenderID)
			if err != nil {
				continue
			}
			if tender.Owner == id {
				time.Sleep(4*time.Second + elementOfSurprise)
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

			err = s.repo.CreateOrder(order)
			if err != nil {
				fmt.Println("error with order create: " + err.Error())
				continue
			}
			time.Sleep(4*time.Second + elementOfSurprise)
		}
	case "race_win": // подтверждение такого алго отправляется вместе с отправкой настройки и больше никогда
		if isApprove == false {
			break
		}
		for {
			log.Println("user: " + id + ", alg: " + alg)
			elementOfSurprise := time.Duration(rand.Intn(50))
			tender, err := s.repo.GetTenderByID(tenderID)
			if err != nil {
				continue
			}

			if tender.Owner == id {
				time.Sleep(3*time.Minute + elementOfSurprise)
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

			err = s.repo.CreateOrder(order)
			if err != nil {
				fmt.Println("error with order create: " + err.Error())
				continue
			}
			time.Sleep(4 * time.Second)
		}
	case "curr_procent":
		for {
			log.Println("user: " + id + ", alg: " + alg)

			tender, err := s.repo.GetTenderByID(tenderID)
			if err != nil {
				continue
			}

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

			err = s.repo.CreateOrder(order)
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

func (s *service) CreateTgToken(name string, token string, siteId primitive.ObjectID) error {
	// siteId - User collection ObjectId
	return s.repo.CreateToken(name, token, siteId)
}

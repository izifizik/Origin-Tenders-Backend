package botService

import (
	"crypto/md5"
	"encoding/binary"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"math/rand"
	"origin-tender-backend/server/internal/repository/mongodb"
	"strconv"
)

type service struct {
	repo mongodb.Repository
}

func NewBotService(repo mongodb.Repository) BotService {
	return &service{repo: repo}
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

func (s *service) StartServeTendor(tendorID primitive.ObjectID) {
	// тут наинает работу бот
	//постоянно (раз в какое то время) мониторить тендер на цену
	// сравнивать цену с минимальной если да то с критической
	// отправлять запросы на подтверждение транзакции
	// транзация висит не на боте а на пользователе

}

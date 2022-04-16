package botService

import (
	"crypto/md5"
	"encoding/binary"
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

// siteId - User collection ObjectId
func (s *service) CreateTgToken(name string, token string, siteId string) error {
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

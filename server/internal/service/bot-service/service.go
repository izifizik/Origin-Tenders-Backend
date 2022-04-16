package botService

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math/rand"
	"strconv"
)

type service struct {
}

func NewBotService() BotService {
	return &service{}
}

func (s *service) GenerateToken(ID string) string {
	h := md5.New()
	io.WriteString(h, "And Leon's getting larger!")
	token := generateToken(binary.BigEndian.Uint64(h.Sum(nil)))

}

func generateToken(seed uint64) string {
	rand.Seed(int64(seed))
	return strconv.Itoa(rand.Int())
}

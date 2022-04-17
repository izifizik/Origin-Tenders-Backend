package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"origin-tender-backend/server/internal/domain"
)

func (r repo) GetTgUsers() ([]domain.TelegramUser, error) {
	var tgUsers []domain.TelegramUser

	cursor, err := r.tgUserCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &tgUsers)
	if err != nil {
		return nil, err
	}

	return tgUsers, nil
}

package main

import (
	"context"
	cryptoRand "crypto/rand"
	"database/sql"
	"debtsapp/internal/service/model"
	"debtsapp/internal/storage"
	"encoding/hex"
	"fmt"
	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const MaxQuantityUsers = 10

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	logger := zap.Must(zap.NewProduction()).Sugar()
	db, err := storage.New()
	if err != nil {
		logger.Fatalw("There was an error trying to configure db", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatalw("Error Trying to close db", err)
		}
	}(db)

	newStorage := storage.InitDB(db)

	for i := 0; i < MaxQuantityUsers; i++ {
		user := model.UserRequest{
			Name:     faker.Name(),
			LastName: faker.LastName(),
			Username: fmt.Sprintf("%s%d", faker.Username(), i),
			Password: faker.Password(),
			Email:    faker.Email(),
		}
		err := newStorage.Users.CreateAndInvite(context.Background(), &user, generateRandomToken())
		if err != nil {
			logger.Fatalw("There was an error trying to create and invite user", err)
			panic(err)
		}
	}
	logger.Infow("Mocked Users created and invited")
}

func generateRandomToken() string {
	bytes := make([]byte, 16)
	if _, err := cryptoRand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

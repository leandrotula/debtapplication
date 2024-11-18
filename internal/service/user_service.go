package service

import (
	"debtsapp/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService interface {
	Save(context *gin.Context)
}

type UserServiceImpl struct {
	storage *storage.Storage
}

func (u *UserServiceImpl) Save(c *gin.Context) {

	var request storage.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error(err)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)
	err := u.storage.Users.Create(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, storage.NewErrorDB(fmt.Sprintf("Error saving user: %s", err.Error())))
	}

	c.JSON(http.StatusCreated, request)
}

func NewUserService(storage *storage.Storage) UserService {
	return &UserServiceImpl{storage: storage}
}

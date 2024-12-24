package service

import (
	customErrors "debtsapp/internal/error"
	model2 "debtsapp/internal/service/model"
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

// RegisterUser godoc
// @Summary     Register User
// @Description Add a new User
// @Tags        user
// @Param       user body model.UserRequest true "User Request Body"
// @Accept      json
// @Produce     json
// @Success     201  {object} model.UserRequest "User Created"
// @Failure     400  {object} any "Bad Request"
// @Failure     500  {object} any "Internal Server Error"
// @Router      /v1/users [post]
func (u *UserServiceImpl) Save(c *gin.Context) {

	var request model2.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error(err)
		customErrors.NewAppError(c, http.StatusBadRequest, "Invalid request")
		return
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

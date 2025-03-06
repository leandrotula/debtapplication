package service

import (
	"debtsapp/internal/configuration"
	customErrors "debtsapp/internal/error"
	model2 "debtsapp/internal/service/model"
	"debtsapp/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService interface {
	CreateAndInvite(context *gin.Context)
	ActivateUser(context *gin.Context)
}

type UserServiceImpl struct {
	configuration *configuration.Application
}

// RegisterUser godoc
// @Summary     Register User and try to send and invitation
// @Description Add a new User in disable state
// @Tags        user
// @Param       user body model.UserRequest true "User Request Body"
// @Accept      json
// @Produce     json
// @Success     201  {object} model.UserRequest "User Created"
// @Failure     400  {object} any "Bad Request"
// @Failure     500  {object} any "Internal Server Error"
// @Router      /v1/users [post]
func (u *UserServiceImpl) CreateAndInvite(c *gin.Context) {

	var request model2.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error(err)
		customErrors.NewAppError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)
	err := u.configuration.Storage.Users.CreateAndInvite(c, &request, uuid.NewString()) //fake token just for now
	if err != nil {
		u.configuration.Logger.Errorw("Couldn't create user", err)
		c.JSON(http.StatusInternalServerError, storage.NewErrorDB(fmt.Sprintf("Error saving user: %s", err.Error())))
	}

	u.configuration.Logger.Infow("User created and invited")
	c.JSON(http.StatusCreated, model2.NewUserResponse(request.Name, request.LastName, request.Username, request.Email))
}

// Activate an user godoc
// @Summary     Activate an user chaning status to activate = true
// @Description Activate a user
// @Tags        user
// @Param       token query string true "Client's token"
// @Produce     json
// @Success     204 "No content"
// @Failure     400  {object} ErrorResponse "Bad Request"
// @Failure     500  {object} ErrorResponse "Internal Server Error"
// @Router      /v1/users [patch]
func (u *UserServiceImpl) ActivateUser(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, NewErrorResponse("token required", http.StatusBadRequest))
		return
	}
	err := u.configuration.Storage.Users.Activate(c, c.Query("token"))
	if err != nil {
		u.configuration.Logger.Errorw("Couldn't activate user", err)
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	c.Status(http.StatusNoContent)
}

func NewUserService(configuration *configuration.Application) UserService {
	return &UserServiceImpl{configuration: configuration}
}

package service

import (
	"debtsapp/internal/configuration"
	customErrors "debtsapp/internal/error"
	clogger "debtsapp/internal/logger"
	"debtsapp/internal/service/encription"
	appmodel "debtsapp/internal/service/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserService interface {
	CreateAndInvite(context *gin.Context)
	ActivateUser(context *gin.Context)
}

type UserServiceImpl struct {
	app               *configuration.Application
	encryptionService encription.CustomEncryption
}

// RegisterUser godoc
// @Summary     Register User and try to send and invitation
// @Description Add a new User in disable state
// @Tags        user
// @Param       user body model.UserRequest true "User Request Body"
// @Accept      json
// @Produce     json
// @Success     201  {object} model.UserResponse "User Created"
// @Failure     400  {object} any "Bad Request"
// @Failure     500  {object} any "Internal Server Error"
// @Router      /v1/users [post]
func (u *UserServiceImpl) CreateAndInvite(c *gin.Context) {

	var request appmodel.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error(err)
		customErrors.NewAppError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	hash, errorEncryption := u.encryptionService.Encrypt([]byte(request.Password), 15)
	if errorEncryption != nil {
		c.JSON(http.StatusInternalServerError, "Error encrypting password")
		return
	}
	request.Password = string(hash)
	err := u.app.Storage.Users.CreateAndInvite(c, &request, uuid.NewString()) //fake token just for now
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error creating user")
		return
	}

	c.JSON(http.StatusCreated, appmodel.NewUserResponse(request.Name, request.LastName, request.Username, request.Email))
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
	logger := clogger.GetLogger()
	token := c.Query("token")
	if token == "" {
		logger.Error("token required")
		c.JSON(http.StatusBadRequest, NewErrorResponse("token required", http.StatusBadRequest))
		return
	}
	err := u.app.Storage.Users.Activate(c, c.Query("token"))
	if err != nil && err.Error() == "user not found" {
		logger.Error(err)
		c.JSON(http.StatusNotFound, NewErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func NewUserService(application *configuration.Application,
	encryption encription.CustomEncryption) UserService {
	return &UserServiceImpl{
		app:               application,
		encryptionService: encryption,
	}
}

package token

import (
	"debtsapp/internal/configuration"
	"github.com/gin-gonic/gin"
)

type AbstractTokenService interface {
	GenerateJwtToken(context *gin.Context)
}

type AppTokenService struct {
	Generator
	configuration *configuration.Application
}

func NewTokenService() AbstractTokenService {
	return &AppTokenService{}
}

// Generates a token for an existing and active user
// @Summary     Generates a token for an existing and active user
// @Description It generates a token
// @Param       user body CreateTokenPayload "Token request body"
// @Accept      json
// @Produce     json
// @Success     201  {object} CreateTokenResponse "User token"
// @Failure     400  {object} any "Bad Request"
// @Failure     500  {object} any "Internal Server Error"
// @Router      /v1/token [post]
func (t *AppTokenService) GenerateJwtToken(context *gin.Context) {
	token, err := t.Generator.GenerateJwtToken("", "", "")
	if err != nil {
		context.Status(500)
	}
	context.JSON(200, token) //will change all this
}

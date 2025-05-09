package token

import (
	"context"
	"debtsapp/internal/configuration"
	customErrors "debtsapp/internal/error"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AbstractTokenService interface {
	GenerateJwtToken(context *gin.Context)
}

type AppTokenService struct {
	Generator
	configuration *configuration.Application
}

func NewTokenService(app *configuration.Application) AbstractTokenService {
	return &AppTokenService{
		Generator:     &CustomToken{},
		configuration: app,
	}
}

// GenerateTokenHandler genera un token JWT para un usuario activo
// @Summary     Generates a token for an existing and active user
// @Description It generates a token
// @Tags        Create User's token
// @Accept      json
// @Produce     json
// @Param       request body CreateTokenPayload true "Token Request Body"
// @Success     201  {object} CreateTokenResponse "User token"
// @Failure     400  {object} any "Bad Request"
// @Failure     500  {object} any "Internal Server Error"
// @Router      /v1/token [post]
func (t *AppTokenService) GenerateJwtToken(c *gin.Context) {
	var tokenPayload CreateTokenPayload
	if err := c.ShouldBindJSON(&tokenPayload); err != nil {
		log.Error(err)
		customErrors.NewAppError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	userFound, err := t.configuration.Storage.Users.FindUserByEmail(context.Background(), tokenPayload.Email)
	if err != nil {
		log.Error(err)
		customErrors.NewAppError(c, http.StatusUnauthorized, "user unauthorized")
	}

	if bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(tokenPayload.Password)) != nil {
		customErrors.NewAppError(c, http.StatusUnauthorized, "user unauthorized")
	}

	claims := jwt.MapClaims{
		"sub": userFound.ID,
		"exp": time.Now().Add(t.configuration.ConfigurationToken.Expiration).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": t.configuration.ConfigurationToken.Issuer,
		"aud": t.configuration.ConfigurationToken.Audience,
	}
	token, err := t.Generator.GenerateJwtToken(t.configuration.ConfigurationToken.Secret, claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewCustomTokenError(err.Error()))
	}
	c.JSON(http.StatusOK, NewTokenResponse(token))
}

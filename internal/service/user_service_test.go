package service

import (
	"debtsapp/internal/configuration"
	"debtsapp/internal/service/encription"
	"debtsapp/internal/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserServiceImpl_CreateAndInvite_error_request_binding(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	invalidJSON := `{"email": "test@example.com"`
	req, err := http.NewRequest(http.MethodPost, "/invite", strings.NewReader(invalidJSON))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	userService := NewUserService(nil, nil)
	userService.CreateAndInvite(c)
	require.Equal(t, http.StatusBadRequest, w.Code)

}

func TestUserServiceImpl_CreateAndInvite_error_encryption(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	invalidJSON := `{
"name": "Juan",
  "last_name": "Pérez",
  "username": "juanperez",
  "password": "testKey",
  "email": "juan.perez@example.com"
}`
	req, err := http.NewRequest(http.MethodPost, "/invite", strings.NewReader(invalidJSON))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	encryptionMock := new(encription.CustomEncryptionMock)
	encryptionMock.On("Encrypt", []byte("testKey"), 15).Return(nil, errors.New("error"))
	userService := NewUserService(nil, encryptionMock)

	userService.CreateAndInvite(c)
	require.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestUserServiceImpl_CreateAndInvite_error_storage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invalidJSON := `{
        "name": "Juan",
        "last_name": "Pérez",
        "username": "juanperez",
        "password": "testKey",
        "email": "juan.perez@example.com"
    }`
	req, err := http.NewRequest(http.MethodPost, "/invite", strings.NewReader(invalidJSON))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	userMock := new(storage.UserMock)
	encryptionMock := new(encription.CustomEncryptionMock)

	encryptionMock.On("Encrypt", mock.Anything, mock.AnythingOfType("int")).Return("encryptedPassword", nil)

	userMock.On("CreateAndInvite",
		mock.AnythingOfType("*gin.Context"),
		mock.AnythingOfType("*model.UserRequest"),
		mock.AnythingOfType("string")).Return(errors.New("error"))

	userService := NewUserService(configuration.NewApplication(
		storage.NewStorage(userMock),
		nil,
	), encryptionMock)

	userService.CreateAndInvite(c)
	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserServiceImpl_CreateAndInvite_ok(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	invalidJSON := `{
        "name": "Juan",
        "last_name": "Pérez",
        "username": "juanperez",
        "password": "testKey",
        "email": "juan.perez@example.com"
    }`
	req, err := http.NewRequest(http.MethodPost, "/invite", strings.NewReader(invalidJSON))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	userMock := new(storage.UserMock)
	encryptionMock := new(encription.CustomEncryptionMock)

	encryptionMock.On("Encrypt", mock.Anything, mock.AnythingOfType("int")).Return("encryptedPassword", nil)

	userMock.On("CreateAndInvite",
		mock.AnythingOfType("*gin.Context"),
		mock.AnythingOfType("*model.UserRequest"),
		mock.AnythingOfType("string")).Return(nil)

	userService := NewUserService(configuration.NewApplication(
		storage.NewStorage(userMock),
		nil,
	), encryptionMock)

	userService.CreateAndInvite(c)
	require.Equal(t, http.StatusCreated, w.Code)
}

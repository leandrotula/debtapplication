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

func TestUserServiceImpl_CreateAndInvite(t *testing.T) {
	t.Run("Test UserServiceImpl_CreateAndInvite_error_request_binding", func(t *testing.T) {
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
	})

	t.Run("Test UserServiceImpl_CreateAndInvite_error_encryption", func(t *testing.T) {

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
	})

	t.Run("Test UserServiceImpl_CreateAndInvite_error_storage", func(t *testing.T) {
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
	})

	t.Run("Test UserServiceImpl_CreateAndInvite_ok", func(t *testing.T) {
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

	})
}

func TestUserServiceImpl_ActivateUser(t *testing.T) {
	t.Run("Test UserServiceImpl_ActivateUser_ok", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, err := http.NewRequest(http.MethodPatch, "/v1/users?token=abcdeftg", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		userMock := new(storage.UserMock)
		encryptionMock := new(encription.CustomEncryptionMock)

		userMock.On("Activate",
			mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("string")).Return(nil)

		userService := NewUserService(configuration.NewApplication(
			storage.NewStorage(userMock),
			nil,
		), encryptionMock)

		userService.ActivateUser(c)
		require.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Test UserServiceImpl_ActivateUser_error_invalid_token", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, err := http.NewRequest(http.MethodPatch, "/v1/users", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		userMock := new(storage.UserMock)
		encryptionMock := new(encription.CustomEncryptionMock)

		userService := NewUserService(configuration.NewApplication(
			storage.NewStorage(userMock),
			nil,
		), encryptionMock)

		userService.ActivateUser(c)
		require.Equal(t, http.StatusBadRequest, w.Code)
		userMock.AssertNumberOfCalls(t, "Activate", 0)

	})

	t.Run("Test UserServiceImpl_ActivateUser_error_activation_storage", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, err := http.NewRequest(http.MethodPatch, "/v1/users?token=somerandomtoken", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		c.Request = req

		userMock := new(storage.UserMock)
		userMock.On("Activate",
			mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("string")).Return(errors.New("user not found"))
		encryptionMock := new(encription.CustomEncryptionMock)

		userService := NewUserService(configuration.NewApplication(
			storage.NewStorage(userMock),
			nil,
		), encryptionMock)

		userService.ActivateUser(c)
		require.Equal(t, http.StatusNotFound, w.Code)
		userMock.AssertNumberOfCalls(t, "Activate", 1)
	})
}

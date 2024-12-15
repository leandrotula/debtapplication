package error

import (
	"github.com/gin-gonic/gin"
)

func NewAppError(c *gin.Context, status int, message string) {
	c.JSON(status, message)
}

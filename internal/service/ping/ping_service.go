package ping

import "github.com/gin-gonic/gin"

// ping verification godoc
// @Summary     ping verify if app is running
// @Description ping operation
// @Tags        ping
// @Produce     json
// @Success     200  {object} string
// @Router      /ping [get]
func Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	}
}

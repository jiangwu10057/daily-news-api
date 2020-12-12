package api

import (
	"news/utils/logging"
	"news/service"

	"github.com/gin-gonic/gin"
)

// Auth 授权
func Auth(c *gin.Context) {
	var service service.AuthService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Auth()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
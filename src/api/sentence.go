package api

import (
	"news/utils/logging"
	"news/service"

	"github.com/gin-gonic/gin"
)

// ListSentence 美句列表接口
func ListSentence(c *gin.Context) {
	service := service.ListSentenceService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowSentence 美句详情接口
func ShowSentence(c *gin.Context) {
	service := service.ShowSentenceService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
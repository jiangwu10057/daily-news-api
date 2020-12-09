package api

import (
	"news/utils/logging"
	"news/service"

	"github.com/gin-gonic/gin"
)

// ListMedia 新闻列表接口
func ListMedia(c *gin.Context) {
	service := service.ListMediaService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowMedia 商品详情接口
func ShowMedia(c *gin.Context) {
	service := service.ShowMediaService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
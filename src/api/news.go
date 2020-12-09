package api

import (
	"news/utils/logging"
	"news/service"

	"github.com/gin-gonic/gin"
)

// ListNews 新闻列表接口
func ListNews(c *gin.Context) {
	service := service.ListNewsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowNews 商品详情接口
func ShowNews(c *gin.Context) {
	service := service.ShowNewsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
package api

import (
	"news/utils/logging"
	"news/service"

	"github.com/gin-gonic/gin"
)

// ShowDailyReport 商品详情接口
func ShowDailyReport(c *gin.Context) {
	service := service.ShowDailyReportService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Error(err)
	}
}
package conf

import (
	"news/cache"
	"news/model"
	"news/utils/logging"
	
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var conf_path = "/www/conf/"

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	envErr := godotenv.Load(conf_path + ".env")

	if envErr != nil {
		logging.Info(envErr)
		panic(envErr)
	}

	gin.SetMode(os.Getenv("GIN_JMODE"))

	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}

	// 读取翻译文件
	if err := LoadLocales(conf_path + "locales/zh-cn.yaml"); err != nil {
		logging.Info(err)
		panic(err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
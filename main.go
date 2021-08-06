package main

import (
	"fmt"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	"scorelist/dao"
	_ "scorelist/docs"
	"scorelist/models"
	"scorelist/routers"
	"scorelist/setting"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://github.com/rickonzhang/scorelist.git
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:9000
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./bubble conf/config.ini")
		return
	}

	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Score{})
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}
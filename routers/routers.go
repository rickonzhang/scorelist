package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"scorelist/controller"
	"scorelist/setting"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Group := r.Group("v1")
	{
		v1Group.POST("/post", controller.CreateScore)
		v1Group.GET("/get/all", controller.GetScorelist)
		v1Group.POST("/update", controller.UpdateAScore)
		v1Group.GET("/delete", controller.DeleteAScore)
		v1Group.GET("/get/single", controller.GetAScore3)
		v1Group.POST("/islike", controller.IsLike)
	}
	return r
}
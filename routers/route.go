package routers

import (
	"goweb/controllers"
	"goweb/db"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//初始化路由
func InitRouter() (router *gin.Engine) {
	router = gin.Default()

	mysql := db.InitMysqlDb()
	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var userC controllers.UserController = &controllers.UserService{SqlDB: mysql}
	user := router.Group("/user")
	{

		user.POST("/regis", userC.Login)
		user.POST("/login", userC.Login)
		user.POST("/logout", userC.Login)
		user.POST("/updateInfo", userC.Login)
	}

	return router
}

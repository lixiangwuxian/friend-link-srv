package webserver

import (
	"github.com/gin-gonic/gin"
	"lxtend.com/friend_link/middleware"
	"lxtend.com/friend_link/webserver/handler"
)

func router(engine *gin.Engine) {
	usr := engine.Group("")
	usr.Use(middleware.CorsMiddleware())
	usr.Use(middleware.RateLimitMiddleware())
	{
		usr.GET("/friend/approve.php", handler.ApproveFriendLink)
		usr.POST("/friend/application.php", handler.ApplicateFriendLink)
		usr.GET("/friend/reject.php", handler.RejectFriendLink)
		usr.GET("/friend/list.json", handler.GetFriendList)
	}

	admin := engine.Group("/admin")
	admin.Use(middleware.CheckAuth())
	{
		admin.DELETE("/friend", handler.DeleteFriendLink)
	}

}

func StartUp() {
	r := gin.Default()
	r.RemoteIPHeaders = []string{"CF-Connecting-IP", "X-Forwarded-For", "X-Real-IP"}
	router(r)
	// 启动Gin服务器
	r.Run(":12881")

}

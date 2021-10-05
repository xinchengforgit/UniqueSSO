package router

import (
	"unique/jedi/controller"
	"unique/jedi/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.TracingMiddleware())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// CAS related
	sr := r.Group("/cas")
	sr.POST("/login", controller.Login)
	sr.POST("/logout", controller.Logout)
	sr.GET("/p3/serviceValidate", controller.ValidateTicket)

	smsrouter := r.Group("/sms")
	smsrouter.POST("code", controller.SendSmsCode)

	qrrouter := r.Group("/qrcode")
	qrrouter.GET("code", controller.GetWorkWxQRCode)

	larkRouter:=r.Group("/lark")
	larkRouter.GET("user",controller.GetLarkUserInfo)//把userInfo打印出来
	larkRouter.GET("/login",controller.LarkLogin)
	r.GET("/oauth/callback/lark",controller.GetAuthCode) //lark登陆后的回调网址
	//redirect到code,并打印用户redirect到/code里面

}

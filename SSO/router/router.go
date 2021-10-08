package router

import (
	"unique/jedi/common"
	"unique/jedi/controller"
	"unique/jedi/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret")) //session加密用的字符串
	r.Use(middleware.TracingMiddleware())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(sessions.Sessions(common.CAS_COOKIE_NAME, store))
	// CAS related
	sr := r.Group("/cas")
	sr.POST("/login", controller.Login)
	sr.POST("/logout", controller.Logout)
	sr.GET("/p3/serviceValidate", controller.ValidateTicket)

	oauth := sr.Group("/oauth")
	oauth.GET("/lark", controller.LoginWithLark)

	smsrouter := r.Group("/sms")
	smsrouter.POST("code", controller.SendSmsCode)

}

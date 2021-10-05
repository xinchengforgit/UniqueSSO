package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"unique/jedi/conf"
	"unique/jedi/util"
)




func LarkLogin(ctx *gin.Context){
	service := ctx.Query("service")
	//redirect主要是从
	url:=util.MakeRedirectUrl(conf.SSOConf.WorkLark.RedirectUri,conf.SSOConf.WorkLark.AppId,service)
	ctx.Redirect(http.StatusMovedPermanently,url)
}

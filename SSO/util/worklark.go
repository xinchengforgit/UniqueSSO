package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type TokenPostBody struct {
	AppId string `json:"app_id"`
	AppSecret string  `json:"app_secret"`
}


type JsonResult struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}
//理清楚步骤
//1.向"https://open.larksuite.com/open-apis/authen/v1/index?redirect_uri=https://sso.hustunique.com/oauth/callback/lark&app_id=cli_a1bf206dffb8d009"
//发送请求,其中Content-Type=application/json
//2.从响应中获取auth_code
//3.向"https://open.larksuite.com/open-apis/auth/v3/app_access_token/internal"发送请求获取token
//4.将携带token和auth_code来获取用户信息"https://open.larksuite.com/open-apis/authen/v1/access_token"
//
//ok

const (
	larkWebFmt="https://open.larksuite.com/open-apis/auth/v3/app_access_token/internal" //获取app_token的网址
	larkStateFmt="https://open.larksuite.com/open-apis/authen/v1/index?redirect_uri=%s&app_id=%s&state=%s"  //获取auth_code的网址
	larkUserIdFmt="https://open.larksuite.com/open-apis/authen/v1/access_token"  //获取用户信息的网址
)


//
//发送post请求,使用json格式
//设计接口,返回值是string,error


func GetLarkAppToken(appId,appSecret string)(string,error){
	fmt.Println("GetLarkAppToken:now GetLarkToken")
	//注意要构建http请求
	//构建一个
	postBody:=&TokenPostBody{AppId: appId,AppSecret: appSecret}
	jsonByte, err := json.Marshal(postBody)
	if err!=nil{
		return "",err
	}
	req, err := http.NewRequest(http.MethodPost, larkWebFmt, bytes.NewBuffer(jsonByte))
	if err!=nil{
		return "",err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8") //设置header
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "",err
	}
	defer response.Body.Close()
	body,_:=ioutil.ReadAll(response.Body)
	var jsonResult JsonResult
	json.Unmarshal(body,&jsonResult)
	//token存入redis内
	//todo
	return jsonResult.AppAccessToken,nil
}

//使用/login?type=lark&service=xxxx登陆
//数据全是从后台发送过来的
//redirect to larkStateFmt这个网址带上redirect_uri app_id state这三个参数
//登陆后浏览器跳转到 /oauth/callback/lark?state=xxx&code=xxx ,并从redis中查询token,如果没有查询到,就获取token
//redirect to  /lark/userInfo 并且打印各种信息,感觉不大行啊,这样的话,有风险啊
//url中的scheme 方案

//使用lark登陆的话,会有一个FetchLarkAuthCode的函数,用于发送请求给lark并返回
//service是之前要使用的服务应用,构建一个request
//不对这个是后台的跳转,并无调用,他会给我重定向吗

func MakeRedirectUrl(redirectUrl,appId ,service string)string{
	//用于构建
	urlStr:=fmt.Sprintf(larkStateFmt,redirectUrl,appId,service)
	return urlStr
	//接下来他就会给我跳转到这个网址
}

//这个的目的是获取

func FetchLarkUserId(authCode,service string){

}

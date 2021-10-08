package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unique/jedi/database"
)

type TokenPostBody struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type JsonResult struct {
	Code           int           `json:"code"`
	Msg            string        `json:"msg"`
	AppAccessToken string        `json:"app_access_token"`
	Expire         time.Duration `json:"expire"`
}

type UserIdPostBody struct {
	Token     string `json:"app_access_token"`
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`
}

type UserInfo struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Name             string `json:"name"`
	EnName           string `json:"en_name"`
	AvatarUrl        string `json:"avatar_url"`
	AvatarThumb      string `json:"avatar_thumb"`
	AvatarMiddle     string `json:"avatar_middle"`
	AvatarBig        string `json:"avatar_big"`
	OpenId           string `json:"open_id"`
	UnionId          string `json:"union_id"`
	Email            string `json:"email"`
	UserId           string `json:"user_id"`
	Mobile           string `json:"mobile"`
	TenantKey        string `json:"tenant_key"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
}
type ResponseInfo struct {
	StatusCode int       `json:"code"`
	Message    string    `json:"msg"`
	UserInfo   *UserInfo `json:"data,omitempty"` //失败时这一项不存在
}

const (
	LarkWebFmt    = "https://open.larksuite.com/open-apis/auth/v3/app_access_token/internal"                  //获取app_token的网址
	LarkStateFmt  = "https://open.larksuite.com/open-apis/authen/v1/index?redirect_uri=%s&app_id=%s&state=%s" //获取auth_code的网址
	LarkUserIdFmt = "https://open.larksuite.com/open-apis/authen/v1/access_token"                             //获取用户信息的网址
)

//获取token并存入redis中

func GetLarkAppToken(appId, appSecret string) (string, error) {
	fmt.Println("GetLarkAppToken:now GetLarkToken")
	//注意要构建http请求
	//构建一个
	postBody := &TokenPostBody{AppId: appId, AppSecret: appSecret}
	jsonByte, err := json.Marshal(postBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, LarkWebFmt, bytes.NewBuffer(jsonByte))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8") //设置header
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var jsonResult JsonResult
	json.Unmarshal(body, &jsonResult)
	//token存入redis内
	//todo
	err = database.RedisClient.Set(context.Background(), "lark_token", jsonResult.AppAccessToken, jsonResult.Expire*time.Second).Err()
	if err != nil {
		return "", err
	}
	return jsonResult.AppAccessToken, nil
}

//只返回open_id其他的可以从response_body里面去拿到

func FetchWorkLarkUserId(token, code string) (string, error) {
	//接下里要构建http请求
	log.Printf("function FetchWorkLarkUserId")
	userIdPostBody := &UserIdPostBody{
		Token:     token,
		GrantType: "authorization_code",
		Code:      code,
	}

	jsonByte, err := json.Marshal(userIdPostBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, LarkUserIdFmt, bytes.NewBuffer(jsonByte))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json") //设置header
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	var responseInfo ResponseInfo
	err = json.Unmarshal(body, responseInfo)
	if err != nil {
		return "", err
	}
	if responseInfo.UserInfo == nil {
		return "", errors.New(responseInfo.Message)
	} //出错的情况
	return responseInfo.UserInfo.OpenId, nil //不能返回response.UserInfo.UserId 是中文
	//返回值可以根据需求修改,详情见上面的结构体

}

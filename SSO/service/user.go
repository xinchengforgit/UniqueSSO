package service

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"net/url"
	"unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/database"
	"unique/jedi/pkg"
	"unique/jedi/util"

	"github.com/sirupsen/logrus"
)

func VerifyUserWithGet(ctx context.Context, data string, signType string) (*database.User, error) {
	switch signType {
	case common.SignTypeLark:
		return VerifyUserByLark(data)
	default:
		return nil, errors.New("Invalid sign type")
	}
}

func VerifyUser(ctx context.Context, login *pkg.LoginUser, signType string) (*database.User, error) {
	switch signType {
	case common.SignTypeEmailPassword:
		return VerifyUserByEmail(login.Email, login.Password)
	case common.SignTypePhonePassword:
		return VerifyUserByPhone(login.Phone, login.Password)
	case common.SignTypePhoneSms:
		return VerifyUserBySMS(ctx, login.Phone, login.Code)
	case common.SignTypeWechat:
		return VerifyUserByQrcode(login.QrcodeSrc)
	case common.SignTypeLark:
		return VerifyUserByLark(login.LarkSrc)
	default:
		return nil, errors.New("Invalid sign type")
	}
}

func GetUserById(uid string) (*database.User, error) {
	user := new(database.User)
	err := database.DB.Table(user.TableName()).Where("uid = ?", uid).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserByEmail(email, password string) (*database.User, error) {
	lg := logrus.WithField("type", "email")
	user := new(database.User)
	err := database.DB.Table(user.TableName()).Where("email = ?", email).Scan(user).Error
	if err != nil {
		lg.WithError(err).Info("get user by email error")
		return nil, err
	}
	if err := util.ValidatePassword(password, user.Password); err != nil {
		lg.WithError(err).Info("validate password error")
		return nil, err
	}
	return user, nil
}

func VerifyUserByPhone(phone, password string) (*database.User, error) {
	user := new(database.User)
	err := database.DB.Table(user.TableName()).Where("phone = ?", phone).Scan(user).Error
	if err != nil {
		return nil, err
	}
	if err := util.ValidatePassword(password, user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserBySMS(ctx context.Context, phone, sms string) (*database.User, error) {
	user := new(database.User)
	code, err := util.GetSMSCodeByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if code != sms {
		return nil, errors.New("sms code is wrong")
	}
	err = database.DB.Table((user.TableName())).Where("phone = ?", phone).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserByQrcode(qrcode string) (*database.User, error) {
	src, err := url.Parse(qrcode)
	if err != nil {
		return nil, err
	}
	code, err := util.FetchAuthCode(src.Query().Get("key"))
	if err != nil {
		return nil, err
	}
	//这一段逻辑要改一下
	//首先得获取app_token()
	//
	conf.SSOConf.WorkWx.AccessToken.RWLock.RLock()
	token := conf.SSOConf.WorkWx.AccessToken.Token
	conf.SSOConf.WorkWx.AccessToken.RWLock.RUnlock()
	//直接利用token和code来获取userid就好了
	//ok
	userid, err := util.FetchWorkwxUserId(token, code)
	if err != nil {
		return nil, err
	}

	user := new(database.User)
	err = database.DB.Table(user.TableName()).Where("workwx_user_id = ?", userid).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

//由于lark重定向回来到的是redirect_uri的页面可以把code作为参数传入/login,故而larkSrcCode实际上是Authorization_code

func VerifyUserByLark(larkSrcCode string) (*database.User, error) {
	//从redis中获取lark_token
	token, err := database.RedisClient.Get(context.Background(), "lark_token").Result()
	if err == redis.Nil {
		//重新获取token并设置
		token, err = util.GetLarkAppToken(conf.SSOConf.Lark.AppId, conf.SSOConf.Lark.AppSecret)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	//这里的唯一表示符是open_id
	userid, err := util.FetchWorkLarkUserId(token, larkSrcCode)
	if err != nil {
		return nil, err
	}

	user := new(database.User)
	err = database.DB.Table(user.TableName()).Where("workLark_user_id = ?", userid).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

package conf

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"regexp"
)

type Conf struct {
	Application  ApplicationConf  `mapstructure:"application"`
	Database     DatabaseConf     `mapstructure:"database"`
	Redis        RedisConf        `mapstructure:"redis"`
	Sms          []SMSOptions     `mapstructure:"sms"`
	OpenPlatform OpenPlatformConf `mapstructure:"openplat_form"`
	APM          APMConf          `mapstructure:"apm"`
	Lark         LarkConf         `mapstructure:"work_lark"` //表示lark的
}
type ApplicationConf struct {
	Host            string           `mapstructure:"host"`
	Port            string           `mapstructure:"port"`
	Name            string           `mapstructure:"name"`
	Mode            string           `mapstructure:"mode"`
	ReadTimeout     int              `mapstructure:"read_timeout"`
	WriteTimeout    int              `mapstructure:"write_timeout"`
	AllowService    []string         `mapstructure:"allow_service"`
	AllowServiceReg []*regexp.Regexp `mapstructure:"-"`
}

type DatabaseConf struct {
	PostgresDSN string `mapstructure:"postgres_dsn"`
}

type RedisConf struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type SMSOptions struct {
	Name       string `mapstructure:"name" validator:"oneof='verificationCode'"`
	TemplateId string `mapstructure:"template_id"`
	SignName   string `mapstructure:"sign_name"`
}

//add Lark Conf

type LarkConf struct {
	AppId       string `mapstructure:"app_id"`
	RedirectUri string `mapstructure:"redirect_uri"`
	AppSecret   string `mapstructure:"app_secret"`
}

type OpenPlatformConf struct {
	GrpcAddr       string `mapstructure:"grpc_addr"`
	GrpcCert       string `mapstructure:"grpc_cert"`
	GrpcServerName string `mapstructure:"grpc_server_name"`
}

type APMConf struct {
	ReporterBackground string `mapstructure:"reporter_backend"`
}

var (
	SSOConf = &Conf{}
)

func InitConf(confFilepath string) error {
	viper.SetConfigFile(confFilepath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(SSOConf)
	if err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(SSOConf); err != nil {
		return err
	}

	SSOConf.Application.AllowServiceReg = make([]*regexp.Regexp, len(SSOConf.Application.AllowService))
	for i, service := range SSOConf.Application.AllowService {
		reg, err := regexp.Compile(service)
		if err != nil {
			return err
		}
		SSOConf.Application.AllowServiceReg[i] = reg
	}

	if SSOConf.Application.Mode == "debug" {
		zapx.Info("run mode", zap.String("mode", SSOConf.Application.Mode))
	}

	return nil
}

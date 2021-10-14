package config

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
*
*   @Title:
*       Config
*
*   @Description:
*		🤠
*
*   @Remarks:
*
*
*   @Functions:
*
*
*   @Author:
*       XiangChen
*
*   @Date:
*       2021/10/8
*
*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  */
import (
	cfg "github.com/XC-Zero/common/config"
)

type Config struct {
	UserDatapackConfig     cfg.MySQLConfiguration
	RedisConfig            cfg.RedisConfiguration
	ESConfig               cfg.ESConfiguration
	MinioConfig            cfg.MinioConfiguration
	IpRealAddressAPIConfig IpRealAddressAPI
}
type IpRealAddressAPI struct {
	URL       string
	AppKey    string
	AppSecret string
	AppCode   string
}

var CONFIG *Config

// InitConfig ...
func InitConfig(configName string, configPaths []string) error {
	CONFIG = new(Config)
	return cfg.InitConfiguration(configName, configPaths, CONFIG)
}

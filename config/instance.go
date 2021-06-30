package config

import "github.com/spf13/viper"

var Instance Config

type Config struct {
	Addr        string
	YouPlusPath bool
	YouPlusAuth bool
	YouPlusUrl  string
}

func ReadConfig() error {
	configer := viper.New()
	configer.AddConfigPath("./")
	configer.AddConfigPath("../")
	configer.SetConfigType("yaml")
	configer.SetConfigName("config")
	err := configer.ReadInConfig()
	if err != nil {
		return err
	}
	configer.SetDefault("addr", ":4300")
	configer.SetDefault("youplus.enablepath", false)
	configer.SetDefault("youplus.url", "http://localhost:8999")
	configer.SetDefault("youplus.auth", false)

	Instance = Config{
		Addr:        configer.GetString("addr"),
		YouPlusPath: configer.GetBool("youplus.enablepath"),
		YouPlusUrl:  configer.GetString("youplus.url"),
		YouPlusAuth: configer.GetBool("youplus.auth"),
	}
	return nil
}

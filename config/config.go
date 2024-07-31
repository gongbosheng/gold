package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	System *SystemConfig `mapstructure:"system" json:"system"`
	Mysql  *MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Logs   *LogsConfig   `mapstructure:"logs" json:"logs"`
}

func InitConfig() error {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用目录失败:%s", err))
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}
	return nil
}

package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Database struct {
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
	DBUsername string `mapstructure:"db_username"`
	DBPassword string `mapstructure:"db_password"`
	DBCharset  string `mapstructure:"db_charset"`
	DBPaseTime string `mapstructure:"db_parse_time"`
	DBLocal    string `mapstructure:"db_local"`
}

type Configuration struct {
	Database `mapstructure:"database"`
}

func initConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file: ", err)
		os.Exit(1)
	}
}

func InitConfigDsn() string {
	initConfig()

	var config Configuration

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("[Config][InitConfigDsn] Unable to decode into struct:", err)
	}

	dsn := config.DBUsername + `:` +
		config.DBPassword + `@tcp(` +
		config.DBHost + `:` +
		config.DBPort + `)/` +
		config.DBName +
		`?charset=` + config.DBCharset +
		`&parseTime=` + config.DBPaseTime +
		`&loc=` + config.DBLocal

	return dsn
}

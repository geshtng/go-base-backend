package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Database struct {
	DBHost            string `mapstructure:"db_host"`
	DBPort            string `mapstructure:"db_port"`
	DBName            string `mapstructure:"db_name"`
	DBUsername        string `mapstructure:"db_username"`
	DBPassword        string `mapstructure:"db_password"`
	DBPostgresSslMode string `mapstructure:"db_postgres_ssl_mode"`
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

	dsn := `postgres://` + config.DBUsername + `:` + config.DBPassword + `@` + config.DBHost + `:` + config.DBPort + `/` + config.DBName + `?sslmode=` + config.DBPostgresSslMode

	return dsn
}

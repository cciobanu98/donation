package config

import (
	"github.com/spf13/viper"
)

// DBConfig handle config items for DB
type DBConfig struct {
	Host 		string
	Port 		int 
	Username 	string
	Password 	string
	Name 		string
	Charset 	string
	Dialect 	string
	DbURI  		string
}

// Config for app
type Config struct {
	DB *DBConfig
}

// Conf is config for app
var Conf Config

// LoadConfig is loading configuration
func LoadConfig() Config{

	viper.SetConfigName("local")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	err := viper.Unmarshal(&Conf)
	if err != nil {
		panic("unable to get config")
	}
	return Conf
}
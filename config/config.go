package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configurations struct {
	Env    string
	Server struct {
		Port string
	}
	JWTKey string
	Ldap   struct {
		Host     string
		Port     string
		User     string
		Password string
		Base     string
	}
}

var C Configurations

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
	// Debug: Print the loaded configuration
	fmt.Printf("Loaded configuration: %+v\n", C)
}

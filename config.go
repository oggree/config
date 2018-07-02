package config

import (
	"github.com/spf13/viper"
	"fmt"
)

var initStatus bool = false

func Construct() bool {

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.AutomaticEnv()
	initStatus = true
	return true
}

func Get(key string) interface{} {
	if !initStatus {
		Construct()
	}

	fmt.Println(key)

	return viper.Get(key)
}

var Config = map[string]interface{}{
	"port": "8081",

	"dbDriver": "mysql",

	"databases": map[string]interface{}{
		"mysql": map[string]interface{}{
			"default": map[string]string{
				"host":"localhost",
				"username":"root",
				"password": "",
				"database":"",
			},
			"blog": map[string]string{
				"host":"localhost",
				"username":"root",
				"password": "",
				"database":"",
			},
		},
	},

	"socialNetworks": map[string]interface{}{
		"facebook": map[string]string{
			"appId": "",
			"appSecret": "",
		},
	},

}
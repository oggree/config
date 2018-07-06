package config

import (
	"github.com/spf13/viper"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

var initStatus bool = false
var configFolder string = "config"

type config struct {
	Port string
	SocialNetworks configSocialNetworks
}

type configSocialNetworks struct {
	Facebook configFacebook
}

type configFacebook struct {
	AppID string
	AppSecret string
}

func Construct() bool {

	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		os.Mkdir(configFolder, 0644)
	}

	emptyConfig := map[string]interface{}{
		"global": config{},
	}

	serializedConfig, configSerializeErr := yaml.Marshal(&emptyConfig)

	//fmt.Printf("--- m dump:\n%s\n\n", string(serializedConfig))
	if configSerializeErr != nil {
		panic(configSerializeErr)
	}

	if _, err := os.Stat(configFolder + "/global.yaml"); os.IsNotExist(err) {
		configCreator("global", serializedConfig)
	}

	viper.SetConfigName("global")
	viper.AddConfigPath("./" + configFolder)

	err1 := viper.ReadInConfig() // Find and read the config file
	if err1 != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err1))
	}

	viper.AutomaticEnv()
	initStatus = true
	return true
}

func InitConfig(fileName string, configDefaults interface{}) bool  {

	emptyConfig := map[string]interface{}{
		"" + fileName: configDefaults,
	}

	serializedConfig, configSerializeErr := yaml.Marshal(&emptyConfig)

	if configSerializeErr != nil {
		panic(configSerializeErr)
	}

	if _, err := os.Stat(configFolder + "/" + fileName + ".yaml"); os.IsNotExist(err) {
		configCreator(fileName, serializedConfig)
	}

	viper.SetConfigName(fileName)
	viper.AddConfigPath("./" + configFolder)

	viper.MergeInConfig()

	return true
}

func Get(key string) interface{} {
	if !initStatus {
		Construct()
	}

	requestedConfig := viper.Get(key)

	if requestedConfig == nil {
		panic("Config key not setted: " + key)
	}

	return requestedConfig
}


func configCreator(configName string, configText []byte) (bool, error) {
	var fPath string = configFolder + "/" + configName + ".yaml"
	//if file doesnt exist
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		generalConfigFile, generalConfigFileCreateErr := os.Create(fPath)
		if generalConfigFileCreateErr!=nil {
			return false, generalConfigFileCreateErr
		}

		defer generalConfigFile.Close()
	}
	//file write
	fileWriteErr := ioutil.WriteFile(fPath, configText, 0644)
	if fileWriteErr!=nil {
		return false, fileWriteErr
	} else {
		return true, nil
	}
}
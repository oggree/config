package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/oggree/logger"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AutomaticEnv()                                   // Automatically use environment variables where available
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace . in config keys with _ in env variables

	files, err := os.ReadDir("./configs")
	if err != nil {
		logger.Error("Error reading config directory", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			ext := filepath.Ext(filename)
			name := strings.TrimSuffix(filename, ext)

			viper.SetConfigName(name)
			viper.SetConfigType(strings.TrimPrefix(ext, "."))
			viper.AddConfigPath("./configs")

			if err := viper.MergeInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					logger.Info("Loaded config file: " + filename)
				} else {
					logger.Error("Error reading config file"+filename, err)
				}
			} else {
				logger.Info("Loaded config file: " + filename)
			}
		}
	}
}

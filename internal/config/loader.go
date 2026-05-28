package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Active *Configuration

func LoadConfiguration(file string) error {
	setupViper(file)
	setDefaults()

	config, err := readConfig()
	if err != nil {
		return err
	}

	Active = config

	return nil
}

func setupViper(file string) {
	viper.SetConfigFile(file)
	viper.SetConfigType("yml")
	viper.SetEnvPrefix("hymetric")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func setDefaults() {
	viper.SetDefault("http.address", "0.0.0.0:8080")
	viper.SetDefault("cache.ttl", "30s")
	viper.SetDefault("cache.cleanupInterval", "45s")
	viper.SetDefault("timers", []Timer{
		{
			Name:        "mytimer",
			DisplayName: "My timer",
			Icon:        "0000",
			Url:         "https://timer.example.com",
		},
	})
}

func readConfig() (*Configuration, error) {
	if err := viper.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	viper.AutomaticEnv()

	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

package config

import (
	"log"
	logging "sbit/infrastructure/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig() {
	readConfig() //read config file

	err := logging.NewLogger(logging.Configuration{
		EnableConsole:     viper.GetBool("logger.console.enable"),
		ConsoleJSONFormat: viper.GetBool("logger.console.json"),
		ConsoleLevel:      viper.GetString("logger.console.level"),
	}, 0)

	if err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}
}
func readConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("reading config error ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config update, automaticarlly reloading services... ")
	})
}

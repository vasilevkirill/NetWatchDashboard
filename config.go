package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Config = viper.New()

func ConfigLoad() {
	checkFatalError(ConfigRead())
	// Наблюдение за изменениями конфигурационного файла
	Config.WatchConfig()
	Config.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		// При изменении конфига перечитываем его
		checkFatalError(ConfigRead())
	})
}
func ConfigRead() error {
	Config.AddConfigPath(".")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	err := Config.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error config file: %s", err.Error())
	}
	return nil
}

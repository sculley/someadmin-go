package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// FileConfig receives the config file path, name, and type
type FileConfig struct {
	Path string
	Name string
	Type string
}

func (c FileConfig) Load(conf interface{}) {
	// Set the file name of the configurations file
	viper.SetConfigName(c.Name)
	// Set the type of the configuration file
	viper.SetConfigType(c.Type)
	// Add the path to look for the configurations file
	viper.AddConfigPath(c.Path)

	log.Debug("Reading config file")

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file, %s", err)
	}

	// Watch for changes
	viper.WatchConfig()
	viper.OnConfigChange(func(evt fsnotify.Event) {
		fmt.Println("Config file changed:", evt.Name)
		err := viper.Unmarshal(conf)
		if err != nil {
			log.Error("Unable to decode into struct, %v", err)
		}
	})

	// Unmarshall the config into the provided struct
	err := viper.Unmarshal(conf)
	if err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}
}

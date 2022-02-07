package viper

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/sculley/someadmin-go/config"
	"github.com/spf13/viper"
)

// FileConfig receiver for config load
type FileConfig struct {
	Name          string
	ConfigType    string
	Folder        string
	WatchCallback func(conf *interface{}, event fsnotify.Event)
}

func (c FileConfig) LoadConfig(conf interface{}) error {
	viper.SetConfigName(c.Name)
	viper.SetConfigType(c.ConfigType)
	viper.AddConfigPath(c.Folder)

	viper.AutomaticEnv()

	// Invoke callback if config changes
	viper.WatchConfig()
	viper.OnConfigChange(func(evt fsnotify.Event) {
		err := viper.Unmarshal(conf)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if c.WatchCallback != nil {
			c.WatchCallback(&conf, evt)
		}
	})

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return &config.ErrConfigNotFound{Err: err, File: c.Folder}
		}

		return &config.ErrInvalidConfig{Err: err}
	}

	// Unmarshall to struct
	err := viper.Unmarshal(conf)
	if err != nil {
		return &config.ErrInvalidConfig{Err: err}
	}

	return nil
}

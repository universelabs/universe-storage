package config

import (
	// universe
	"github.com/universelabs/universe-server/universe"
	// deps
	"github.com/spf13/viper"
)

// This holds the configuration constants
type Constants struct {
	Port string
	StormDB struct {
		Path string
	}
}

// Reads from config file and populates a Constants struct
func NewConstants() (*Constants, error) {
	// config filename and path
	viper.AddConfigPath("./config")
	// this is currently hardcoded
	viper.SetConfigName("server")
	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}
	// unmarshal from config file
	consts := &Constants{}	
	err = viper.Unmarshal(consts)
	return consts, err
}
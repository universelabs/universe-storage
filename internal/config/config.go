package config

import (
	"github.com/spf13/viper"
	"github.com/asdine/storm"
)

type Constants struct {
	PORT string
	StormDB struct {
		Path string
	}
}

func initViper() (Constants, error) {
	// config filename and path
	viper.SetConfigName("./configs/universe.config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}
	// default PORT value
	viper.SetDefault("port", "8080")
	// unmarshal from config file
	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}

type Config struct {
	Constants
	KS storage.Keystore
}

// generates a new configuration instance to be passed around 
func New() (*Config, error) {
	cfg := Config{}
	constants, cfgerr := initViper()
	cfg.Constants = constants
	if cfgerr != nil {
		return &cfg, cfgerr
	}
	cfg.KS = storage.Keystore{}
	// need to check if creating new db or opening existing one
	dberr := cfg.KS.Init(cfg.Constants.StormDB.Path)
	if dberr != nil {
		cfg.KS = nil
		return &cfg, dberr
	}
	return &cfg, nil
}
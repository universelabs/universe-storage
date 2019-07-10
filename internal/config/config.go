package config

import (
	"github.com/universelabs/universe-server/storage"

	"github.com/spf13/viper"
)

type Constants struct {
	Port string
	StormDB struct {
		Path string
		exists bool
	}
}

func initViper() (Constants, error) {
	// config filename and path
	viper.AddConfigPath("./config")
	viper.SetConfigNmae("server")
	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}
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
	dberr := cfg.KS.Init(cfg.Constants.StormDB.Path, !cfg.StormDB.exists)
	if !cfg.StormDB.exists {
		viper.Set("StormDB.exists", true)
	}
	if dberr != nil {
		cfg.KS = nil
		return &cfg, dberr
	}
	return &cfg, nil
}
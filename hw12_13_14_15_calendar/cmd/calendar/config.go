package main

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
[logger]
level = "DEBUG"
`)

func InitConfig() error {
	if configFile != "" {
		return initByFile(configFile)
	}
	return initDefault()
}

func initByFile(cf string) error {
	cfDir, cfName := path.Split(cf)
	viper.SetConfigName(strings.Split(cfName, ".")[0])
	viper.SetConfigType("toml")

	viper.AddConfigPath(cfDir)
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		var e viper.ConfigFileNotFoundError
		if errors.As(err, &e) {
			fmt.Printf("%v. Initing default configuration.\n", err)
			return initDefault()
		}
		return fmt.Errorf("error by reading config file %w", err)
	}
	return nil
}

func initDefault() error {
	return viper.ReadConfig(bytes.NewBuffer(defaultConfig))
}

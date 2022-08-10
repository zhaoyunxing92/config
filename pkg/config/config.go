package config

import (
	"github.com/spf13/viper"
)

type Kind uint

const (
	Map Kind = iota
	Slice
	Struct
)

type Config struct {
	Name string
	Path string
}

func Load(path string) error {
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

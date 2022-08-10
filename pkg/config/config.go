package config

import (
	"encoding/json"

	"github.com/spf13/viper"

	"config/pkg/extension"
)

func Load(path string) error {
	var (
		err  error
		data []byte
	)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	for _, conf := range extension.GetConfigs() {
		switch conf.Kind() {
		case extension.Map:
			for _, v := range viper.GetStringMap(conf.Prefix()) {

				if data, err = json.Marshal(v); err != nil {
					return err
				}
				if err = json.Unmarshal(data, conf); err != nil {
					return err
				}
				if err = conf.Load(); err != nil {
					return err
				}
			}
		case extension.Struct:
			if err = viper.UnmarshalKey(conf.Prefix(), conf); err != nil {
				return err
			}
			if err = conf.Load(); err != nil {
				return err
			}
		}
	}
	return nil
}

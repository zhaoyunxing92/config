package config

import (
	"config/pkg/extension"
)

type Dubbo struct {
	Application Application         `json:"application"`
	Provider    Provider            `json:"provider"`
	Registries  map[string]Registry `json:"registries"`
}

func (d *Dubbo) Prefix() string {
	return "dubbo"
}

func (d *Dubbo) Load() error {
	return nil
}

func (d *Dubbo) Order() int {
	return 0
}

func (d *Dubbo) Kind() extension.Kind {
	return extension.Struct
}

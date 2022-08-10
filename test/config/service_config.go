package config

import (
	"config/pkg/extension"
)

// Service is the configuration of the service provider
type Service struct {
	Filter    string          `yaml:"filter" json:"filter,omitempty" property:"filter"`
	Interface string          `yaml:"interface"  json:"interface,omitempty" property:"interface"`
	Cluster   string          `default:"failover" yaml:"cluster"  json:"cluster,omitempty" property:"cluster"`
	Version   string          `yaml:"version"  json:"version,omitempty" property:"version" `
	Methods   []*MethodConfig `yaml:"methods"  json:"methods,omitempty" property:"methods"`
	Retries   string          `yaml:"retries"  json:"retries,omitempty" property:"retries"`
}

func (sc *Service) Kind() extension.Kind {
	return extension.Map
}

func (sc *Service) Prefix() string {
	return "dubbo.provider.services"
}

func (sc *Service) Load() error {
	return nil
}

func (sc *Service) Order() int {
	return 5
}

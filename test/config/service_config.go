package config

import "config/pkg/config"

// ServiceConfig is the configuration of the service provider
type ServiceConfig struct {
	Filter    string          `yaml:"filter" json:"filter,omitempty" property:"filter"`
	Interface string          `yaml:"interface"  json:"interface,omitempty" property:"interface"`
	Cluster   string          `default:"failover" yaml:"cluster"  json:"cluster,omitempty" property:"cluster"`
	Version   string          `yaml:"version"  json:"version,omitempty" property:"version" `
	Methods   []*MethodConfig `yaml:"methods"  json:"methods,omitempty" property:"methods"`
	Retries   string          `yaml:"retries"  json:"retries,omitempty" property:"retries"`
}

func (sc *ServiceConfig) Kind() config.Kind {
	return config.Map
}

func (sc *ServiceConfig) Prefix() string {
	return "dubbo.provider.services"
}

func (sc *ServiceConfig) Load() error {
	//TODO implement me
	panic("implement me")
}

func (sc *ServiceConfig) Order() int {
	return 5
}

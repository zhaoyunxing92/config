package config

// Service is the configuration of the service provider
type Service struct {
	Filter    string          `yaml:"filter" json:"filter,omitempty" property:"filter"`
	Interface string          `yaml:"interface"  json:"interface,omitempty" property:"interface"`
	Cluster   string          `default:"failover" yaml:"cluster"  json:"cluster,omitempty" property:"cluster"`
	Version   string          `yaml:"version"  json:"version,omitempty" property:"version" `
	Methods   []*MethodConfig `yaml:"methods"  json:"methods,omitempty" property:"methods"`
	Retries   string          `yaml:"retries"  json:"retries,omitempty" property:"retries"`
}

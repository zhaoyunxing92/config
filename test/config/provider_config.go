package config

// Provider is the default configuration of service provider
type Provider struct {
	Filter string `yaml:"filter" json:"filter,omitempty" property:"filter"`
	// Deprecated Register whether registration is required
	Register bool `yaml:"register" json:"register" property:"register"`
	// Registry is registry ids list
	Registry []string `yaml:"registry" json:"registry" property:"registry"`
	// Services services
	Services     map[string]*Service `yaml:"services" json:"services,omitempty" property:"services"`
	ProxyFactory string              `default:"default" yaml:"proxy" json:"proxy,omitempty" property:"proxy"`
	// adaptive service
	AdaptiveService        bool `default:"false" yaml:"adaptive-service" json:"adaptive-service" property:"adaptive-service"`
	AdaptiveServiceVerbose bool `default:"false" yaml:"adaptive-service-verbose" json:"adaptive-service-verbose" property:"adaptive-service-verbose"`
}

package config

// MethodConfig defines method config
type MethodConfig struct {
	Name        string `yaml:"name"  json:"name,omitempty" property:"name"`
	Retries     string `yaml:"retries"  json:"retries,omitempty" property:"retries"`
	LoadBalance string `yaml:"loadbalance"  json:"loadbalance,omitempty" property:"loadbalance"`
	Weight      int64  `yaml:"weight"  json:"weight,omitempty" property:"weight"`
}

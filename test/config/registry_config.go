package config

// Registry is the configuration of the registry center
type Registry struct {
	Protocol  string `validate:"required" yaml:"protocol"  json:"protocol,omitempty" property:"protocol"`
	Timeout   string `default:"5s" validate:"required" yaml:"timeout" json:"timeout,omitempty" property:"timeout"` // unit: second
	Group     string `yaml:"group" json:"group,omitempty" property:"group"`
	Namespace string `yaml:"namespace" json:"namespace,omitempty" property:"namespace"`
	TTL       string `default:"10s" yaml:"ttl" json:"ttl,omitempty" property:"ttl"` // unit: minute
	Address   string `validate:"required" yaml:"address" json:"address,omitempty" property:"address"`
	Zone      string `yaml:"zone" json:"zone,omitempty" property:"zone"`       // The region where the registry belongs, usually used to isolate traffics
	Weight    int64  `yaml:"weight" json:"weight,omitempty" property:"weight"` // Affects traffic distribution among registriesConfig, useful when subscribe to multiple registriesConfig Take effect only when no preferred registry is specified.
}

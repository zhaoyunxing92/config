package config

type Dubbo struct {
	Application Application         `json:"application"`
	Provider    Provider            `json:"provider"`
	Registries  map[string]Registry `json:"registries"`
}

func (d *Dubbo) Prefix() string {
	return "dubbo"
}

func (d *Dubbo) Process(config map[string]interface{}) error {
	return nil
}

func (d *Dubbo) Order() int {
	return 0
}

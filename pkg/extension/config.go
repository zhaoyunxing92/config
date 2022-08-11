package extension

import (
	"log"
	"sort"
	"sync"
)

var (
	configs map[string]Config
	lock    sync.Mutex
)

type Config interface {
	Prefix() string

	Process(config map[string]interface{}) error

	Order() int
}

// ConfigPostProcess config post process
type ConfigPostProcess interface {
	Process(map[string]interface{})
}

func Register(name string, config Config) {
	lock.Lock()
	defer lock.Unlock()
	if configs != nil {
		_, found := configs[name]
		if found {
			log.Fatalf("config plugin %s was registered twice", name)
		}
	} else {
		configs = map[string]Config{}
	}

	configs[name] = config
}

func GetConfigs() []Config {
	lock.Lock()
	defer lock.Unlock()
	var cs []Config
	for _, cfg := range configs {
		cs = append(cs, cfg)
	}
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Order() < cs[j].Order()
	})
	return cs
}

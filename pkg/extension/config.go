package extension

import (
	"log"
	"sort"
	"sync"

	"config/pkg/config"
)

var (
	configs map[string]Config
	lock    sync.Mutex
)

type Config interface {
	Prefix() string

	Load() error

	Order() int

	Kind() config.Kind
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
	for _, config := range configs {
		cs = append(cs, config)
	}
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Order() < cs[j].Order()
	})
	return cs
}

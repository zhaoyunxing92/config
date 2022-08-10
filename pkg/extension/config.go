package extension

import (
	"log"
	"sort"
	"sync"
)

type Kind uint

const (
	Map Kind = iota
	Slice
	Struct
)

var (
	configs map[string]Config
	lock    sync.Mutex
)

type Config interface {
	Prefix() string

	Load() error

	Order() int

	Kind() Kind
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

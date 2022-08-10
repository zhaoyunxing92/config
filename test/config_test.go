package test

import (
	"testing"

	"config/pkg/config"
	"config/pkg/extension"
	cfg "config/test/config"
)

func TestLoad(t *testing.T) {
	extension.Register("dubbo", &cfg.Dubbo{})
	//extension.Register("application", &config.Application{})
	//extension.Register("provider", &config.Provider{})
	//extension.Register("registry", &config.Registry{})
	//extension.Register("service", &config.Service{})
	err := config.Load("../configs")
	if err != nil {
		t.Fatalf("config load err:%v", err)
	}
}

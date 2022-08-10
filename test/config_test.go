package test

import (
	"testing"

	"config/pkg/extension"
	"config/test/config"
)

func TestLoad(t *testing.T) {
	extension.Register("application", &config.Application{})
	extension.Register("provider", &config.ProviderConfig{})
	extension.Register("registry", &config.RegistryConfig{})
	extension.Register("registry", &config.RegistryConfig{})
}

package config

import (
	"config/pkg/extension"
)

type Application struct {
	Organization string `default:"dubbo-go" yaml:"organization" json:"organization,omitempty" property:"organization"`
	Name         string `default:"dubbo.io" yaml:"name" json:"name,omitempty" property:"name"`
	Module       string `default:"sample" yaml:"module" json:"module,omitempty" property:"module"`
	Version      string `default:"3.0.0" yaml:"version" json:"version,omitempty" property:"version"`
	Owner        string `default:"dubbo-go" yaml:"owner" json:"owner,omitempty" property:"owner"`
	Environment  string `default:"dev" yaml:"environment" json:"environment,omitempty" property:"environment"`

	MetadataType string `default:"local" yaml:"metadata-type" json:"metadataType,omitempty" property:"metadataType"`
}

func (a *Application) Prefix() string {
	return "dubbo.application"
}

func (a *Application) Load() error {
	return nil
}

func (a *Application) Order() int {
	return 1
}

func (a *Application) Kind() extension.Kind {
	return extension.Struct
}

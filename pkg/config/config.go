package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/rawbytes"

	"config/pkg/extension"
)

func Load(path string) error {
	var (
		err   error
		bytes []byte
		koan  *koanf.Koanf
	)

	koan = koanf.New(".")
	if bytes, err = os.ReadFile(path); err != nil {
		log.Fatal(err)
	}

	if err = koan.Load(rawbytes.Provider(bytes), yaml.Parser()); err != nil {
		log.Fatal(err)
	}

	for _, key := range koan.Keys() {
		value := koan.Get(key)
		if str, ok := value.(string); ok && needAnalysis(str) {
			str = str[2 : len(str)-1]
			split := strings.Split(str, ":")
			value = analysis(8, split[0], strings.Join(split[1:], ""), koan)
			k2 := koanf.New(".")
			//kv
			if err = k2.Load(confmap.Provider(map[string]interface{}{key: value}, "."), nil); err != nil {
				fmt.Println("meg=", err)
			}
			_ = koan.Merge(k2)
		}
	}
	for _, conf := range extension.GetConfigs() {
		if err = koan.UnmarshalWithConf(conf.Prefix(), conf, koanf.UnmarshalConf{Tag: "yaml"}); err != nil {
			return err
		}
		if err = conf.Process(koan.All()); err != nil {
			return err
		}
	}
	return nil
}

func analysis(deep int, path, dv string, koan *koanf.Koanf) interface{} {
	if deep <= 0 || !koan.Exists(path) {
		return dv
	}
	value := koan.Get(path)
	if str, ok := value.(string); ok && needAnalysis(str) {
		str = str[2 : len(str)-1]
		split := strings.Split(str, ":")
		deep--
		analysis(deep, split[0], strings.Join(split[1:], ""), koan)
	}
	return value
}

func needAnalysis(str string) bool {
	return strings.HasPrefix(str, "${") && strings.HasSuffix(str, "}")
}

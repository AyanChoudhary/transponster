package utils

import (
	"log"
	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

// GenerateConfigReader creates and a returns a parsed config instance
func GenerateConfigReader(configPath string) *koanf.Koanf {
	k := koanf.New(".")
	extension := filepath.Ext(configPath)

	switch extension {
	case ".json":
		if err := k.Load(file.Provider(configPath), json.Parser()); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
		break

	case ".yaml":
		if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
		break
	}

	return k
}

package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"os"
	"path"
)

type Config struct {
	Port     int
	LogLevel string `default:"warn"`
}

func New() (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	conf := aconfig.Config{
		AllowUnknownFields: true,
		SkipFlags:          true,
		Files:              []string{path.Join(dir, ".env")},
		FileDecoders: map[string]aconfig.FileDecoder{
			".env": aconfigdotenv.New(),
		},
	}

	cfg := &Config{}
	loader := aconfig.LoaderFor(cfg, conf)
	if err = loader.Load(); err != nil {
		return nil, err
	}

	return cfg, nil
}

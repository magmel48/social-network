package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"os"
	"path"
)

type MySQL struct {
	Port     int    `required:"true"`
	Database string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
}

type Config struct {
	Port     int    `default:"3000"`
	LogLevel string `default:"warn"`
	MySQL    MySQL  `env:"MYSQL"`
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

package config

import (
	"learn-go/errors"
)

var DefaultConfig = []byte(`
logger:
  level: "info"

listen: ":8888"

prefix: "/ajay-verse"

is_prod_mode: false

mongo:
  meta_uri: "mongodb+srv://ajay:RjEGri696X2Rw8dK@my-cluster.0uuox.mongodb.net/?retryWrites=true&w=majority&appName=My-Cluster"
`)

type Config struct {
	Logger     Logger `koanf:"logger"`
	Listen     string `koanf:"listen"`
	Prefix     string `koanf:"prefix"`
	IsProdMode bool   `koanf:"is_prod_mode"`
	Mongo      Mongo  `koanf:"mongo"`
}

type Logger struct {
	Level string `koanf:"level"`
}

type Mongo struct {
	MetaURI string `koanf:"meta_uri"`
}

// Validate validates the configuration
func (c *Config) Validate() error {
	ve := errors.ValidationErrs()

	if c.Listen == "" {
		ve.Add("listen", "cannot be empty")
	}
	if c.Logger.Level == "" {
		ve.Add("logger.level", "cannot be empty")
	}

	if c.Mongo.MetaURI == "" {
		ve.Add("mongo.meta_uri", "cannot be empty")
	}

	return ve.Err()
}

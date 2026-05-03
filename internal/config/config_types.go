package config

import (
	"fmt"
	"strings"
)

type Config struct {
	Port         string `env:"PORT" required:"false" default:"8080"`
	Debug        string `env:"DEBUG" required:"false" default:"false"`
	Env          string `env:"ENV" required:"false" default:"dev"`
	LogLevel     string `env:"LOG_LEVEL" required:"false" default:"info"`
	Database_URL string `env:"DB_URL" required:"true"`
}

type EnvValidationError struct {
	Fields  []string
	Message string
}

func (e *EnvValidationError) Error() string {
	return fmt.Sprintf("%s: [%s]", e.Message, strings.Join(e.Fields, ", "))
}

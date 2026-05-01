package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string `env:"PORT" required:"false" default:"8080"`
	Debug string `env:"DEBUG" required:"false" default:"false"`
}

type EnvValidationError struct {
	Fields  []string
	Message string
}

func (e *EnvValidationError) Error() string {
	return fmt.Sprintf("%s: [%s]", e.Message, strings.Join(e.Fields, ", "))
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	var missing []string

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)

		envKey := structField.Tag.Get("env")
		if envKey == "" {
			continue
		}

		value := os.Getenv(envKey)

		if value == "" {
			value = structField.Tag.Get("default")
		}

		field.SetString(value)

		if structField.Tag.Get("required") == "true" && value == "" {
			missing = append(missing, envKey)
		}
	}

	if len(missing) > 0 {
		return nil, &EnvValidationError{
			Fields:  missing,
			Message: "missing required env variables",
		}
	}

	return cfg, nil
}

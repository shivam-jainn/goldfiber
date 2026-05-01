package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func LoadConfig(cfg interface{}) error {
	_ = godotenv.Load()

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
		return &EnvValidationError{
			Fields:  missing,
			Message: "missing required env variables",
		}
	}

	return nil
}

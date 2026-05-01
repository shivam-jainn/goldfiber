package config

import (
	"os"
	"testing"
)

type TestConfig struct {
	Port   string `env:"PORT" required:"false" default:"8080"`
	APIKey string `env:"API_KEY" required:"true"`
}

func TestLoadConfig(t *testing.T) {
	testCases := []struct {
		name        string
		envVars     map[string]string
		expected    Config
		expectError bool
	}{
		{
			name: "all env vars set",
			envVars: map[string]string{
				"PORT":         "3000",
				"DEBUG":        "true",
				"ENV":          "prod",
				"LOG_LEVEL":    "debug",
				"DATABASE_URL": "postgres://user:pass@localhost/db",
			},
			expected: Config{
				Port:     "3000",
				Debug:    "true",
				Env:      "prod",
				LogLevel: "debug",
			},
			expectError: false,
		},
		{
			name: "no env vars, use defaults except required",
			envVars: map[string]string{
				"DATABASE_URL": "postgres://user:pass@localhost/db",
			},
			expected: Config{
				Port:     "8080",
				Debug:    "false",
				Env:      "dev",
				LogLevel: "info",
			},
			expectError: false,
		},
		{
			name: "some env vars set",
			envVars: map[string]string{
				"PORT":         "4000",
				"DEBUG":        "true",
				"DATABASE_URL": "postgres://user:pass@localhost/db",
			},
			expected: Config{
				Port:     "4000",
				Debug:    "true",
				Env:      "dev",
				LogLevel: "info",
			},
			expectError: false,
		},
		{
			name: "missing required field",
			envVars: map[string]string{
				"PORT":  "4000",
				"DEBUG": "true",
			},
			expected:    Config{},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Unset all env vars first
			allKeys := []string{"PORT", "DEBUG", "ENV", "LOG_LEVEL", "DATABASE_URL"}
			for _, k := range allKeys {
				os.Unsetenv(k)
			}

			// Set env vars for this test case
			for k, v := range tc.envVars {
				os.Setenv(k, v)
			}

			cfg := &Config{}
			err := LoadConfig(cfg)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if _, ok := err.(*EnvValidationError); !ok {
					t.Errorf("expected EnvValidationError, got %T", err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if cfg == nil {
					t.Errorf("expected config, got nil")
				} else if *cfg != tc.expected {
					t.Errorf("expected %+v, got %+v", tc.expected, *cfg)
				}
			}
		})
	}
}

func TestLoadConfigWithCustomStruct(t *testing.T) {
	os.Unsetenv("PORT")
	os.Setenv("API_KEY", "secret")
	defer func() {
		os.Unsetenv("API_KEY")
		os.Unsetenv("PORT")
	}()

	cfg := &TestConfig{}
	err := LoadConfig(cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := TestConfig{Port: "8080", APIKey: "secret"}
	if *cfg != expected {
		t.Errorf("expected %+v, got %+v", expected, *cfg)
	}
}

func TestEnvValidationError(t *testing.T) {
	err := &EnvValidationError{
		Fields:  []string{"DATABASE_URL", "API_KEY"},
		Message: "missing required env variables",
	}
	expected := "missing required env variables: [DATABASE_URL, API_KEY]"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}

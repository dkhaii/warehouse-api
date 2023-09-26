package config

import (
	"testing"
)

type mockConfig struct {
	values map[string]string
}

func (c mockConfig) Get(key string) string {
	return c.values[key]
}

func TestNewMySQLDatabase(t *testing.T) {
	type testCase struct {
		name          string
		config        Config
		expectedError bool
	}

	testCases := []testCase{
		{
			name: "Valid Configuration",
			config: mockConfig{
				values: map[string]string{
					"DB_USERNAME": "root",
					"DB_PASSWORD": "",
					"DB_HOST":     "127.0.0.1",
					"DB_PORT":     "3306",
					"DB_DATABASE": "cozy_warehouse",
				},
			},
			expectedError: false,
		},
		{
			name: "Invalid Port",
			config: mockConfig{
				values: map[string]string{
					"DB_USERNAME": "root",
					"DB_PASSWORD": "",
					"DB_HOST":     "127.0.0.1",
					"DB_PORT":     "3306",
					"DB_DATABASE": "cozy_warehouse",
				},
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dbConn, err := NewMySQLDatabase(tc.config)

			if tc.expectedError && err == nil {
				t.Error("Expected an error, but got nil")
			}

			if !tc.expectedError && err == nil {
				t.Errorf("Expected no error, but got an error: %v", err)
			}

			if dbConn != nil {
				defer dbConn.Close()
			}
		})
	}
}

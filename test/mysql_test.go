package test

import (
	"strconv"
	"testing"

	"github.com/dkhaii/warehouse-api/config"
)

type mockConfig struct {
	values map[string]string
}

func (c mockConfig) GetString(key string) string {
	return c.values[key]
}

func (c mockConfig) GetInt(key string) int {
	// Implement logic untuk mengubah string ke int sesuai kebutuhan Anda
	// Contoh:
	value, err := strconv.Atoi(c.values[key])
	if err != nil {
		return 0 // atau nilai default lainnya
	}
	return value
	// return 0
}

func (c mockConfig) GetCloudENV(key string) string {
    // Mengembalikan nilai default berdasarkan key jika diperlukan
    switch key {
    case "PORT":
        return "8080" // Nilai port default untuk pengujian
    case "DB_USERNAME":
        return "testuser" // Nilai username database default untuk pengujian
    default:
        return "" // Mengembalikan string kosong untuk kunci yang tidak diatur dalam pengujian
    }
}

func TestNewMySQLDatabase(t *testing.T) {
	type testCase struct {
		name          string
		config        config.Config
		expectedError bool
	}

	testCases := []testCase{
		{
			name: "Valid Configuration",
			config: mockConfig{
				values: map[string]string{
					"DB_USERNAME": "root",
					"DB_PASSWORD": "",
					"DB_HOST":     "0.0.0.0",
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
					"DB_HOST":     "0.0.0.0",
					"DB_PORT":     "3306",
					"DB_DATABASE": "cozy_warehouse",
				},
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dbConn, err := config.NewMySQLDatabase(tc.config)

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

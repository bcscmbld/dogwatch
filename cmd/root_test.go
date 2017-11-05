package cmd_test

import (
	"os"
	"testing"

	"github.com/gugahoi/dogwatch/cmd"
)

func TestCheckAPI(t *testing.T) {
	testCases := []struct {
		desc      string
		name      string
		api       string
		apiEnv    string
		shouldErr bool
	}{
		{
			desc:      "version no api",
			name:      "version",
			api:       "",
			apiEnv:    "",
			shouldErr: false,
		},
		{
			desc:      "version api flag",
			name:      "version",
			api:       "api-from-flag",
			apiEnv:    "",
			shouldErr: false,
		},
		{
			desc:      "list no api",
			name:      "list",
			api:       "",
			apiEnv:    "",
			shouldErr: true,
		},
		{
			desc:      "list api from flag",
			name:      "list",
			api:       "api-from-flag",
			apiEnv:    "",
			shouldErr: false,
		},
		{
			desc:      "list api from env",
			name:      "list",
			api:       "",
			apiEnv:    "some-env-api",
			shouldErr: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			key := "DOGNZB_API"

			// restore env variable for after the test
			if value := os.Getenv(key); value != "" {
				defer func() { os.Setenv(key, value) }()
			}

			// set the env to the value in the test
			os.Setenv(key, tC.apiEnv)

			// act
			err := cmd.CheckAPI(tC.name, &tC.api)

			// assert
			if tC.shouldErr {
				if err == nil {
					t.Errorf("expected err to not be '%v', got '%v'", nil, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected err to be '%v', got '%v'", nil, err)
				}
			}
		})
	}
}

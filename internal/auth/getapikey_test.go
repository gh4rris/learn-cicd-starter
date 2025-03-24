package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected string
		expErr   string
	}{
		{
			expErr: "o authorization header included",
		},
		{
			key:    "Authorization",
			value:  "-",
			expErr: "malformed authorization header",
		},
		{
			key:    "Authorization",
			value:  "Bearer 12345",
			expErr: "malformed authorization header",
		},
		{
			key:      "Authorization",
			value:    "ApiKey 12245",
			expected: "12345",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case: #%v", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)
			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey: %v", err)
				return
			}
			if output != test.expected {
				t.Errorf("Unexpected: TestGetAPIKey: %v", output)
				return
			}
		})
	}

}

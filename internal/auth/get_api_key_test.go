package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		want      string
		expectErr string
	}{
		{
			name:      "missing authorization header",
			expectErr: "no authorization header",
		},
		{
			name:      "empty authorization header",
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			name:      "malformed authorization header",
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
		{
			name:      "incorrect authorization scheme",
			key:       "Authorization",
			value:     "Bearer xxxxxx",
			expectErr: "malformed authorization header",
		},
		{
			name:  "valid API key",
			key:   "Authorization",
			value: "ApiKey xxxxxx",
			want:  "xxxxxx",
		},
	}
	return t.Fatalf("no tests to run")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headers := http.Header{}
			if test.key != "" {
				headers.Set(test.key, test.value)
			}

			got, err := GetAPIKey(headers)
			if test.expectErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", test.expectErr)
				}
				if !strings.Contains(err.Error(), test.expectErr) {
					t.Fatalf("expected error containing %q, got %q", test.expectErr, err.Error())
				}
				return 
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != test.want {
				t.Errorf("expected: %q, got: %q", test.want, got)
			}
		})
	}
}

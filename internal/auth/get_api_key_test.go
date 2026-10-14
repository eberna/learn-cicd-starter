package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey() returned an unexpected error: %v", err)
	}

	want := "test-key"
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

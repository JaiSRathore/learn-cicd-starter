package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("valid api key", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey my-secret-key")

		key, err := GetAPIKey(headers)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if key != "my-secret-key" {
			t.Fatalf("expected my-secret-key, got %s", key)
		}
	})

	t.Run("missing authorization header", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)

		if err == nil {
			t.Fatal("expected an error")
		}
	})
}

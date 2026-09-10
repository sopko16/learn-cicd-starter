package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string
		authHeader string
		expectedKey string
		expectErr bool
	}{
		{
			name: "valid API key",
			authHeader: "ApiKey abc123",
			expectedKey: "abc123",
			expectErr: false,
		},
		{
			name: "missing authorization header",
			authHeader: "",
			expectedKey: "",
			expectErr: true,
		},
		{
			name: "wrong auth type",
			authHeader: "Bearer abc123",
			expectedKey: "",
			expectErr: true,
		},
		{
			name: "missing API key",
			authHeader: "Apikey",
			expectedKey: "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			key, err := GetAPIKey(headers)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expecteed an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpcted error: %v", err)
			}

			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}

func TestGetAPIKeyNoHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
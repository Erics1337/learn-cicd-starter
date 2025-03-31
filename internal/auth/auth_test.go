package auth

import (
	"errors" // Added import
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name: "Valid ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey mysecretkey"},
			},
			expectedKey: "mysecretkey",
			expectedErr: nil,
		},
		{
			name:        "No Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Header - Wrong Prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer mysecretkey"},
			},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed Header - Missing Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey "},
			},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header: key is missing"),
		},
		{
			name: "Malformed Header - No Space",
			headers: http.Header{
				"Authorization": []string{"ApiKeymysecretkey"},
			},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed Header - Only ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}

			// Compare error messages if both are non-nil
			if err != nil && tt.expectedErr != nil {
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error %q, got %q", tt.expectedErr, err)
				}
			} else if err != tt.expectedErr { // Handles cases where one is nil and the other isn't
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

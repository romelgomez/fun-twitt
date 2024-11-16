package helper_test

import (
	"strings"
	"testing"

	"funtwitt/helper"
)

func TestGeneratePrefixedID(t *testing.T) {
	tests := []struct {
		prefix    string
		minLength int
		expectErr bool
	}{
		{"user_", 20, false},
		{"tweet_", 25, false},
		{"follow_", 30, false},
		{"short_", 5, true}, // Invalid case: minLength < prefix length
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			id, err := helper.GeneratePrefixedID(tt.prefix, tt.minLength)
			if (err != nil) != tt.expectErr {
				t.Errorf("unexpected error status: got %v, want %v", err != nil, tt.expectErr)
			}

			if err == nil {
				if !strings.HasPrefix(id, tt.prefix) {
					t.Errorf("ID does not have expected prefix: got %s, want prefix %s", id, tt.prefix)
				}
				if len(id) < tt.minLength {
					t.Errorf("ID length is less than expected: got %d, want at least %d", len(id), tt.minLength)
				}
			}
		})
	}
}

func TestGeneratePrefixedID_Unique(t *testing.T) {
	idSet := make(map[string]struct{})
	const iterations = 1000

	for i := 0; i < iterations; i++ {
		id, err := helper.GeneratePrefixedID("test_", 20)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, exists := idSet[id]; exists {
			t.Fatalf("duplicate ID generated: %s", id)
		}
		idSet[id] = struct{}{}
	}
}

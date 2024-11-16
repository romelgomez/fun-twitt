package helper_test

import (
	"funtwitt/helper"
	"testing"
	"time"
)

func TestDateTimeToStr(t *testing.T) {
	tests := []struct {
		name           string
		input          *time.Time
		expectedOutput *string
	}{
		{
			name:           "Nil input",
			input:          nil,
			expectedOutput: nil,
		},
		{
			name: "Valid time input",
			input: func() *time.Time {
				t := time.Date(2023, 11, 15, 10, 0, 0, 0, time.UTC)
				return &t
			}(),
			expectedOutput: func() *string {
				s := "2023-11-15T10:00:00+00:00"
				return &s
			}(),
		},
		{
			name: "Edge case: Start of epoch",
			input: func() *time.Time {
				t := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
				return &t
			}(),
			expectedOutput: func() *string {
				s := "1970-01-01T00:00:00+00:00"
				return &s
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.DateTimeToStr(tt.input)
			if (got == nil) != (tt.expectedOutput == nil) {
				t.Errorf("Expected output %v, but got %v", tt.expectedOutput, got)
				return
			}

			if got != nil && *got != *tt.expectedOutput {
				t.Errorf("Expected output %s, but got %s", *tt.expectedOutput, *got)
			}
		})
	}
}

func TestDateTimeToStr_Memory(t *testing.T) {
	t.Run("Pointer Memory Test", func(t *testing.T) {
		tm := time.Now()
		output := helper.DateTimeToStr(&tm)

		if output == nil {
			t.Fatalf("Expected non-nil output, got nil")
		}

		expected := tm.UTC().Format("2006-01-02T15:04:05-07:00")
		if *output != expected {
			t.Errorf("Expected output %s, but got %s", expected, *output)
		}
	})
}

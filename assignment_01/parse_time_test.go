package assignment01_test

import (
	parsetime "assignment/assignment_01/parse_time"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		name        string
		timeStr     string
		expected    time.Time
		time        time.Time
		expectError bool
	}{
		{
			name:        "Valid time string",
			timeStr:     time.RFC850,
			time:        time.Date(2023, time.February, 1, 9, 15, 30, 0, time.UTC),
			expected:    time.Date(2023, time.February, 1, 9, 15, 30, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Invalid time string",
			timeStr:     time.RFC850,
			time:        time.Now(),
			expected:    time.Time{},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parsedTime, err := parsetime.ParseTime(test.timeStr, test.time.String())
			if test.expectError {
				if err == nil {
					t.Error("Expected an error but got nil")
				}
				return
			}
			if !parsedTime.Equal(test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, parsedTime)
			}
		})
	}
}

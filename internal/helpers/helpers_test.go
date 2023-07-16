package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "should return empty string",
			text:     "",
			expected: "",
		},
		{
			name:     "should return BDEFGA",
			text:     "ABDEFGABEF",
			expected: "ABDEFG",
		},
		{
			name:     "should return B",
			text:     "BBBB",
			expected: "B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestSubstring(tt.text)
			assert.Equal(t, tt.expected, got)
		})
	}
}

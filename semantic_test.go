package ptpp_test

import (
	"testing"

	"gopkg.in/ptpp.v1"

	. "github.com/stretchr/testify/assert"
)

func TestDefaultSemanticMatcher(t *testing.T) {
	var semanticMatcher ptpp.DefaultSemanticMatcher
	semanticMatcher.Train("best", "quality")

	tests := []struct {
		context     string
		suggestions []string
		best        string
		matched     bool
	}{
		{"best", []string{"quantity", "quality"}, "quality", true},
		{"worst", []string{"quantity", "quality"}, "quantity", false},
	}
	for _, tt := range tests {
		t.Run(tt.context, func(t *testing.T) {
			best, matched := semanticMatcher.Match(tt.context, tt.suggestions)
			Equal(t, tt.best, best)
			Equal(t, tt.matched, matched)
		})
	}
}

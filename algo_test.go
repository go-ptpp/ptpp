package ptpp_test

import (
	"testing"

	"gopkg.in/ptpp.v1"

	. "github.com/stretchr/testify/assert"
)

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		name string
		v    string
		w    string
		want int
	}{
		{"Insert", "bad", "band", 1},
		{"Delete", "bind", "bid", 1},
		{"Replace", "bulk", "bill", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ptpp.Levenshtein(tt.v, tt.w)
			Equal(t, tt.want, got)
		})
	}
}

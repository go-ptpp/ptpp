package ptpp_test

import (
	"testing"

	"gopkg.in/ptpp.v1"

	. "github.com/stretchr/testify/assert"
)

func TestDefaultSpellChecker(t *testing.T) {
	var spellChecker ptpp.DefaultSpellChecker
	spellChecker.Train([]string{"quality", "quantity", "quantify"})

	tests := []struct {
		word string
		want []string
	}{
		{"quantity", []string{"quantity", "quantify"}},
		{"quanlity", []string{"quality", "quantity"}},
		{"unknown", []string{"unknown"}},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := spellChecker.Check(tt.word)
			ElementsMatch(t, got, tt.want)
		})
	}
}

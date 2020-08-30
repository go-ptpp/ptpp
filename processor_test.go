package ptpp_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"gopkg.in/ptpp.v1"

	. "github.com/stretchr/testify/assert"
)

func TestProcessor(t *testing.T) {
	var processor ptpp.Processor

	processor.Train([]string{
		"bass guitar",
		"spanish rosetta stone",
	})

	tests := []struct {
		phrase string
		want   []string
	}{
		{"electric base guitarr", []string{"electric", "bass guitar"}},
		{"english rosetta stone", []string{"english", "rosetta stone"}},
		{"spannish rosetta stone", []string{"spanish rosetta stone"}},
	}
	for _, tt := range tests {
		t.Run(tt.phrase, func(t *testing.T) {
			got, _ := processor.Process(strings.NewReader(tt.phrase))
			ElementsMatch(t, got, tt.want)
		})
	}

	workingDir, err := os.Getwd()
	if !NoError(t, err) {
		return
	}

	filePath := path.Join(workingDir, "test.zip")
	defer os.Remove(filePath)

	if !NoError(t, processor.Save(filePath)) {
		return
	}

	NoError(t, processor.Load(filePath))
}

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFileSize(t *testing.T) {
	tests := []struct {
		Input  string
		Output int64
	}{
		{"1", 1},
		{"1k", 1 * 1024},
		{"1K", 1 * 1024},
		{"1KB", 1 * 1024},
		{"1M", 1 * 1024 * 1024},
		{"1G", 1 * 1024 * 1024 * 1024},
	}
	for i, v := range tests {
		n, err := ParseFileSize(v.Input)
		assert.Nil(t, err, fmt.Sprintf("#%d", i))
		assert.Equal(t, v.Output, n, fmt.Sprintf("#%d", i))
	}
}

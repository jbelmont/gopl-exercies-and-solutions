package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	actual := echo([]string{"arg1", "arg2", "arg3", "arg4"})
	var expected = "arg1 arg2 arg3 arg4"
	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestEcho2(t *testing.T) {
	args := []string{"arg1", "arg2", "arg3", "arg4"}
	idx, actual := echo2(args)
	assert.Equal(t, len(idx), 4)
	assert.NotNil(t, actual)
}

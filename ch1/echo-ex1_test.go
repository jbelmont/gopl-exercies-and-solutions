package main

import "testing"

func TestEcho(t *testing.T) {
	actual := echo([]string{"arg1", "arg2", "arg3", "arg4"})
	var expected = "arg1 arg2 arg3 arg4"
	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

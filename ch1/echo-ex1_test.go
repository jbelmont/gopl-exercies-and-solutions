package main

import (
	"testing"
	"time"

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

/* Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
 * Section 1.6 illustrates part of the time package
 * Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
 */
func TestTimeForEcho1(t *testing.T) {
	start := time.Now()
	echo([]string{"arg1", "arg2", "arg3", "arg4"})
	assert.NotNil(t, time.Since(start).Seconds())
}

func TestTimeForEcho2(t *testing.T) {
	start := time.Now()
	echo2([]string{"arg1", "arg2", "arg3", "arg4"})
	assert.NotNil(t, time.Since(start).Seconds())
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo([]string{"arg1", "arg2", "arg3", "arg4"})
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2([]string{"arg1", "arg2", "arg3", "arg4"})
	}
}

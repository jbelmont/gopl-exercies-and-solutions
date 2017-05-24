package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	numbers := [...]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	actual := reverse(&numbers)
	expected := [10]int{
		9, 8, 7, 6, 5, 4, 3, 2, 1,
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("%v should equal %v", actual, expected)
		}
	}
}

package maps

import (
	"testing"
)

func TestGetKey(t *testing.T) {
	listOfCities := make(map[string]string, 10)
	listOfCities["north carolina"] = "raleigh"

	actual := getKey(listOfCities, "north carolina")
	expected := "raleigh"
	if actual != expected {
		t.Error("should equal raleigh")
	}
}

package math

import "testing"

const DEFAULT_ERROR = "Expected %v, received %v"

func TestAverage(t *testing.T) {
	expected := 7.28
	result := Average(7.2, 9.9, 6.1, 5.9)

	if result != expected {
		t.Errorf(DEFAULT_ERROR, expected, result)
	}
}

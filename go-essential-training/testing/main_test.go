package main

import (
	"errors"
	"fmt"
	"testing"
)

// One-level structs could be compared directly
func TestExerciseCorrectPattern(t *testing.T) {
	actual, err := exercisePattern("AAA")

	expected := AAA{
		A_1: "Arrange",
		A_2: "Act",
		A_3: "Assert",
	}

	if err != nil {
		t.Errorf("Test case error: %v", err)
	}

	if actual != expected {
		t.Errorf("Actual: %v; Expected: %v", actual, expected)
	}
}

func TestExerciseWrongPattern(t *testing.T) {
	actual, err := exercisePattern("DRY")

	expected := AAA{}

	if errors.Is(err, fmt.Errorf("error: wrong pattern DRY")) {
		t.Errorf("Test case error: %v", err)
	}

	if actual != expected {
		t.Errorf("Actual: %v; Expected: %v", actual, expected)
	}
}

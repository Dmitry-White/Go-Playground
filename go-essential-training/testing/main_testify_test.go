package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExerciseCorrectPatternTestify(t *testing.T) {
	actual, err := exercisePattern("AAA")

	expected := AAA{
		A_1: "Arrange",
		A_2: "Act",
		A_3: "Assert",
	}

	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func TestExerciseWrongPatternTestify(t *testing.T) {
	actual, err := exercisePattern("DRY")

	expected := AAA{}
	expectedError := "error: wrong pattern DRY"

	assert.EqualError(t, err, expectedError)

	assert.Equal(t, expected, actual)
}

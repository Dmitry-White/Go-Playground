package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVerbsTestify(t *testing.T) {
	testCases := []TestCase{
		{"A_1", "Arrange"},
		{"A_2", "Act"},
		{"A_3", "Assert"},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("should get a correct verb by %s key", testCase.value)

		t.Run(testName, func(t *testing.T) {
			actual := getVerb(testCase.value)

			assert.Equal(t, testCase.expected, actual)
		})
	}
}

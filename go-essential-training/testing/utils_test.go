package main

import (
	"fmt"
	"testing"
)

func TestGetVerbs(t *testing.T) {
	testCases := []TestCase{
		{"A_1", "Arrange"},
		{"A_2", "Act"},
		{"A_3", "Assert"},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("should get a correct verb by %s key", testCase.value)

		t.Run(testName, func(t *testing.T) {
			actual := getVerb(testCase.value)

			if actual != testCase.expected {
				t.Errorf("Actual %s; Expected %s", actual, testCase.expected)
			}
		})
	}
}

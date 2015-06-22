package testutil

import "fmt"

// Checks whether two slices contain the same elements.
func ShouldSetEqual(actual interface{}, expected ...interface{}) (msg string) {
	actualSlice, expectedSlice := actual.([]string), expected[0].([]string)

	msg = fmt.Sprintf("Expected:\t%q\nActual:\t%q\n(Should have the same elements)", expectedSlice, actualSlice)

	if len(actualSlice) != len(expectedSlice) {
		return
	}

	actualSet := make(map[string]bool)
	for _, a := range actualSlice {
		actualSet[a] = true
	}

	for _, e := range expectedSlice {
		if actualSet[e] != true {
			return
		}
	}

	expectedSet := make(map[string]bool)
	for _, e := range expectedSlice {
		expectedSet[e] = true
	}

	for _, a := range actualSlice {
		if expectedSet[a] != true {
			return
		}
	}

	return ""
}

func ShouldConsistOf(actual interface{}, expected ...interface{}) (msg string) {
	actualSlice, expectedSlice := actual.([]string), []string{}

	for _, v := range expected {
		expectedSlice = append(expectedSlice, v.(string))
	}

	msg = fmt.Sprintf("Expected:\t%q\nActual:\t%q\n(Should consist of the given elements", expectedSlice, actualSlice)

	if len(actualSlice) != len(expectedSlice) {
		return
	}

	for i, v := range actualSlice {
		if v != expectedSlice[i] {
			return
		}
	}

	return ""
}
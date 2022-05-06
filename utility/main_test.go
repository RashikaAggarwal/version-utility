package main

import (
	"os"
	"testing"
	"errors"
	"reflect"
)

//Function to test main method
func TestMainProgram(t *testing.T) {
	//Case 1: Happy Path
	os.Args = []string{"", "1.2.9.9", "1.3.4"}
	main()
}

//Function to test compareVersions method 
func TestCompareVersions(t *testing.T){
	//Case 1: Happy Path version1 greater than version 2
	input1 := []string{"1.3.4", "1.3"}
	expectedOutput := 1
	actualOutput, _ := compareVersions(input1[0], input1[1])
	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:1 Passed")
	} else {
		t.Errorf("Unexpected Result, got %v wanted %v", actualOutput, expectedOutput)
	}

	//Case 2: Happy Path version1 less than version 2
	input2 := []string{"1.1.4", "1.3"}
	expectedOutput = -1
	actualOutput, _ = compareVersions(input2[0], input2[1])
	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:2 Passed")
	} else {
		t.Errorf("Unexpected Result, got %v wanted %v", actualOutput, expectedOutput)
	}

	//Case 3: Happy Path version1 equals version 2
	input3 := []string{"1.3", "1.3"}
	expectedOutput = 0
	actualOutput, _ = compareVersions(input3[0], input3[1])
	if reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Log("Test Case:3 Passed")
	} else {
		t.Errorf("Unexpected Result, got %v wanted %v", actualOutput, expectedOutput)
	}

	//Case 4: Error in case of invalid version1
	input4 := []string{"test", "1.3"}
	expectedOutput1 := errors.New("Error: Invalid input version")
	_, actualOutput1 := compareVersions(input4[0], input4[1])
	if reflect.DeepEqual(actualOutput1, expectedOutput1) {
		t.Log("Test Case:4 Passed")
	} else {
		t.Errorf("Unexpected Result, got %v wanted %v", actualOutput1, expectedOutput1)
	}

	//Case 5: Error in case of invalid version2
	input5 := []string{"1.1", "test"}
	expectedOutput1 = errors.New("Error: Invalid input version")
	_, actualOutput1 = compareVersions(input5[0], input5[1])
	if reflect.DeepEqual(actualOutput1, expectedOutput1) {
		t.Log("Test Case:5 Passed")
	} else {
		t.Errorf("Unexpected Result, got %v wanted %v", actualOutput1, expectedOutput1)
	}


}
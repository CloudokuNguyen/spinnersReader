package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	expected := "./spinners.json"

	actual := getJsonPath("   ")

	if actual != expected {
		t.Errorf("Test Failed. Expected %s. Got %s", expected, actual)
	}
}

func TestConvertToStringArray(t *testing.T){

	expected := []string{"1","2","3"}

	var interfaceValues []interface{} = make([]interface{}, len(expected))

	for i, _ := range expected {
		interfaceValues[i] = i+1
	}

	actual := convertToStringArray(interfaceValues)

	for i,_ := range actual{
		if expected[i] != actual[i]{
			t.Errorf("Test Failed. Expected %s. Got %s", expected[i], actual[i])
		}
	}

}

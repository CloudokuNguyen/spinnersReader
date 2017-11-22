package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"time"
	"flag"
	"strings"
)

func printStringsWithInterval(interval float64, pattern []string) {

	for _, e := range pattern {

		duration := time.Duration(interval) * time.Millisecond

		time.Sleep(time.Duration(duration))

		fmt.Println(e)

	}

}

func convertToStringArray(t []interface{}) []string {

	s := make([]string, len(t))
	for i, v := range t {
		s[i] = fmt.Sprint(v)
	}

	return s

}

//If path is nil, return the spinners.json file in the same directory
func getJsonPath(path string) string {

	path = strings.TrimSpace(path)

	if path == "" {

		path = "./spinners.json"
	}

	return path
}

func printSpinner(jsonPath string, spinnerPattern string) {

	path := getJsonPath(jsonPath)
	theJson, _ := ioutil.ReadFile(path)

	var data map[string]interface{}
	json.Unmarshal([]byte(theJson), &data)

	spinnerData := data[spinnerPattern].(map[string]interface{})
	interval := spinnerData["interval"].(float64)
	frames := spinnerData["frames"].([]interface{})
	f := convertToStringArray(frames)

	printStringsWithInterval(interval, f)

}

func main() {

	jsonPath := flag.String("path", "", "Path of the json file. If argument is nil, return the path of the spinners.json file in the same directory")
	spinnerPattern := flag.String("pattern", "", "Use one of the pattern in the .json file")

	flag.Parse()

	printSpinner(*jsonPath, *spinnerPattern)

}

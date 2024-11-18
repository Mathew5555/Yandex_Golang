package main

import (
	"encoding/json"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student
	json.Unmarshal(jsonData, &students)
	for i, _ := range students {
		students[i].Grade += 1
	}
	jsonBytes, err := json.Marshal(students)
	return jsonBytes, err
}

func main() {
	modifyJSON([]byte(`[
			{
				"name": "Oleg",
				"grade": 12
			}
		]`))
}

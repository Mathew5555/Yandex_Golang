package main

import "encoding/json"

type Student3 struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var students []Student3
	ans := make(map[string][]Student3)
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return map[string][]byte{}, err
	}
	for _, el := range students {
		if _, ok := ans[el.Class]; ok {
			ans[el.Class] = append(ans[el.Class], el)
			continue
		}
		ans[el.Class] = []Student3{el}
	}

	res := make(map[string][]byte)
	for key, val := range ans {
		res[key], _ = json.Marshal(val)
	}
	return res, nil
}

func main() {
	splitJSONByClass([]byte(`[
        {
            "name": "Oleg",
            "class": "9B"
        },
        {
            "name": "Ivan",
            "class": "9A"
        },
        {
            "name": "Maria",
            "class": "9B"
        },
        {
            "name": "John",
            "class": "9A"
        }
    ]`))
}

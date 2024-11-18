package main

import "encoding/json"

type Student2 struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var ans []Student2
	for _, el := range jsonDataList {
		var ls []Student2
		err := json.Unmarshal(el, &ls)
		if err != nil {
			return []byte{}, err
		}
		ans = append(ans, ls...)
	}
	return json.Marshal(ans)
}

func main() {
	mergeJSONData([]byte(`[
			{
				"name": "Oleg",
				"class": "9B"
			},
			{
				"name": "Ivan",
				"class": "9A"
			}
		]`))
}

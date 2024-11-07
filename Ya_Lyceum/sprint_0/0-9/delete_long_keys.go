package main

func DeleteLongKeys(m map[string]int) map[string]int {
	res := make(map[string]int)
	for key, val := range m {
		if len(key) >= 6 {
			res[key] = val
		}
	}
	return res
}

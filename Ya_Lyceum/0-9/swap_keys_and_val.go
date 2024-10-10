package main

import "fmt"

func SwapKeysAndValues(m map[string]string) map[string]string {
	res := make(map[string]string)
	for key, val := range m {
		res[val] = key
	}
	return res
}

func main() {
	m := map[string]string{
		"Яндекс": "74957397000", "Музей Яндекса": "74991101133",
	}
	fmt.Println(SwapKeysAndValues(m))
}

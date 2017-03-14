package main

import "fmt"

func main() {
	mapofstuff := map[string]string{
		"Tim":    "Good morning",
		"Jenny":  "Bonjour",
		"Harold": "Huzzar",
	}
	fmt.Println(mapofstuff)

	if _, present := mapofstuff["Jerry"]; present {
		fmt.Println(present)
	} else {
		fmt.Println(present)
	}

	for key, value := range mapofstuff {
		fmt.Println(key, " - ", value)
	}
}

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	fmt.Println(m)

	m = map[string]string{
		"name":  "fasil",
		"age":   "18",
		"place": "nilambur",
	}
	fmt.Println(m)

	for i := 0; i < 3; i++ {
		fmt.Println(m)
	}
}

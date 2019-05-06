package main

import ("fmt"
		"strings"
		"os")

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println()
	for _,e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
}

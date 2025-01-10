package main

import (
	"fmt"
	"strings"
)

// var Age int = 20
// Name := "aaaa"

func MakeSuffix(suffix string) func(string) string {
	return func(str string) string {
		if strings.HasSuffix(str, suffix) {
			return str
		}
		return str + suffix
	}
}

func main() {
	f := MakeSuffix(".jpg")
	fmt.Println(f("photo01"))
	fmt.Println(f("photo02.jpg"))
}

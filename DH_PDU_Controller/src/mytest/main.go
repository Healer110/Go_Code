package main

import (
	"fmt"
)

func main() {
	var f string

	for {
		fmt.Print("Input string: ")
		fmt.Scanf("%s\n", &f)
		fmt.Println(f)
	}

}

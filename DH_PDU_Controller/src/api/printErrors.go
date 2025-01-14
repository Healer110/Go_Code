package api

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintErrors(err string) {
	fmt.Printf("%s\n\n", color.RedString("Error: ", err))
}

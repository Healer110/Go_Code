package api

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintMsg(info string) {
	fmt.Printf("%s\n", color.GreenString(info))
}

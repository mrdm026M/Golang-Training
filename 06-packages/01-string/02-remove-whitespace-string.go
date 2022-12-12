// Remove or Strip all white spaces from a string in Go
package main

import (
	"fmt"
	"strings"
)

func main() {
	sample := " This is a sample string   "
	noSpaceString := strings.ReplaceAll(sample, " ", "")
	fmt.Println(noSpaceString)
}

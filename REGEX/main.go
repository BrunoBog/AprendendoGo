package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`(?m)^\d{0,2}(\.\d{1,2})?$`)
	var str = `32,65
                36,256
                1658984,88
                
                `

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}

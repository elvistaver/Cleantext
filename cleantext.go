package main

import (
	"fmt"
	"strings"
)

func CleanText(text string) string {

	if text == "" {
		return ""
	}
	filter := []rune{}
	input1 := strings.ToUpper(text)

	for _, ch := range input1 {

		if ch < 'A' || ch > 'Z' {
			continue
		}
		filter = append(filter, ch)
	}
	return string(filter)
}
func main() {
	fmt.Println(CleanText("45the 1 boy 234 is here"))
}

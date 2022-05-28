package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

var permutations []string
var str []rune

func generateParenthesis(n int) []string {
	permutations = []string{}
	str = []rune{}
	for i := 0; i < 2*n; i++ {
		str = append(str, ' ')
	}
	recursiveGenerate(0, 0, 0, n, 0)
	return permutations
}
func recursiveGenerate(currentlyOpen, currentlyclosed, usedOpen, maxOpen, index int) {

	if currentlyclosed == maxOpen {
		permutations = append(permutations, string(str))
		return
	}
	if currentlyOpen == 0 && usedOpen < maxOpen {
		str[index] = '('
		recursiveGenerate(currentlyOpen+1, currentlyclosed, usedOpen+1, maxOpen, index+1)
		return
	}
	if usedOpen < maxOpen {
		str[index] = '('
		recursiveGenerate(currentlyOpen+1, currentlyclosed, usedOpen+1, maxOpen, index+1)
		str[index] = ')'
		recursiveGenerate(currentlyOpen-1, currentlyclosed+1, usedOpen, maxOpen, index+1)
		return
	}
	str[index] = ')'
	recursiveGenerate(currentlyOpen-1, currentlyclosed+1, usedOpen, maxOpen, index+1)
}

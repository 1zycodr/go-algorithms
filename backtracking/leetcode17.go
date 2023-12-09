package main

import "fmt"

var letters = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	backtrack(&res, []byte{}, 0, digits)
	return res
}

func backtrack(
	res *[]string,
	curr []byte,
	currIdx int,
	digits string,
) {
	if currIdx == len(digits) {
		*res = append(*res, string(curr))
		return
	}
	currLetters := letters[digits[currIdx]]
	for i := 0; i < len(currLetters); i++ {
		curr = append(curr, currLetters[i])
		backtrack(res, curr, currIdx+1, digits)
		curr = curr[:len(curr)-1]
	}
}

func main() {
	res := letterCombinations("22")
	for _, v := range res {
		fmt.Println(v)
	}
}

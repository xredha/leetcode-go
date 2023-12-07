package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Example 1:

Input: num = "52"
Output: "5"
Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.

Example 2:

Input: num = "4206"
Output: ""
Explanation: There are no odd numbers in "4206".

Example 3:

Input: num = "35427"
Output: "35427"
Explanation: "35427" is already an odd number.

*/

func main() {
	var input = []string{"52", "4206", "35427"}

	fmt.Println(largestOddNumber(input[0]))
	fmt.Println(largestOddNumber(input[1]))
	fmt.Println(largestOddNumber(input[2]))
}

func largestOddNumber(num string) string {
	var result string
	strToArr := strings.Split(num, "")
	for i := len(strToArr) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(strToArr[i])
		if x%2 != 0 {
			result = num[:(i + 1)]
			break
		}
	}

	return result
}

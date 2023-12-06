package main

import "fmt"

/*
Example 1:

Input: n = 4
Output: 10
Explanation: After the 4th day, the total is 1 + 2 + 3 + 4 = 10.

Example 2:

Input: n = 10
Output: 37
Explanation: After the 10th day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4) = 37. Notice that on the 2nd Monday, Hercy only puts in $2.

Example 3:

Input: n = 20
Output: 96
Explanation: After the 20th day, the total is (1 + 2 + 3 + 4 + 5 + 6 + 7) + (2 + 3 + 4 + 5 + 6 + 7 + 8) + (3 + 4 + 5 + 6 + 7 + 8) = 96.

*/

func main() {
	var input = []int{4, 10, 20}

	fmt.Println(totalMoney(input[0]))
	fmt.Println(totalMoney(input[1]))
	fmt.Println(totalMoney(input[2]))
}

func totalMoney(n int) int {
	result := 0
	week := 0
	for i := 1; i <= n; i++ {
		result = result + i + week - week*7
		if i%7 == 0 {
			week++
		}
	}
	return result
}

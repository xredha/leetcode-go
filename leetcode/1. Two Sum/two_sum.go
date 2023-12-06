package main

import "fmt"

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func main() {
	var nums = [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}
	var target = []int{9, 6, 6}

	fmt.Println(twoSum(nums[0], target[0]))
	fmt.Println(twoSum(nums[1], target[1]))
	fmt.Println(twoSum(nums[2], target[2]))
}

func twoSum(nums []int, target int) []int {
	var result = []int{}
	for i := 0; i < len(nums); i++ {
		for j := 1 + i; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
	}
	return result
}

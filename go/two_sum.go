package main

import (
	"fmt"
)

// THIS WAY IS THE BEST WAY WHEN ARRAY IS UNSORTED, BUT HAVE ANOTHER WAY WHEN ARRAY IS SORTED, WITH POINTERS
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for index, element := range nums {
		if idx, present := hash[element]; present {
			return []int{idx, index}
		}

		hash[target-element] = index
	}

	return []int{}
}

func main() {
	examples := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9}, // → [0 1]
		{[]int{3, 2, 4}, 6},      // → [1 2]
		{[]int{3, 3}, 6},         // → [0 1]
	}

	for i, ex := range examples {
		result := twoSum(ex.nums, ex.target)
		fmt.Printf("example %d: input nums = %v, target = %d → output: %v\n", i+1, ex.nums, ex.target, result)
	}
}

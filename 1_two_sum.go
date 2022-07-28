package main

func twoSum(nums []int, target int) []int {
	valToIdx := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if idx, ok := valToIdx[complement]; ok {
			return []int{idx, i}
		}

		valToIdx[v] = i
	}

	return nil
}

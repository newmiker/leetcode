package main

func removeElement(nums []int, val int) int {
	tail := len(nums)

	for i := 0; i < tail; i++ {
		if nums[i] != val {
			continue
		}

		for j := tail - 1; ; j-- {
			if j == i {
				// Current element should be removed. There are no elements to swap instead.
				// So returns the index which is len(filtered)-1.
				return j
			}
			if nums[j] == val {
				continue
			}
			nums[i] = nums[j]
			tail = j
			break
		}
	}

	return tail
}

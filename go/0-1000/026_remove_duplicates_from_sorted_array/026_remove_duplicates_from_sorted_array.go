package removeduplicates

func removeDuplicates(nums []int) int {
	prev := nums[0]
	length := 1

	for i := 1; i < len(nums); i++ {
		if prev == nums[i] {
			continue
		}
		prev = nums[i]
		nums[length] = prev
		length++
	}

	return length
}

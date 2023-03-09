package removeelement

func removeElement(nums []int, val int) int {
	length := 0

	for i, v := range nums {
		if v == val {
			continue
		}
		nums[length] = nums[i]
		length++
	}

	return length
}

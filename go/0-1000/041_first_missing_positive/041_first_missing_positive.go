package firstmissingpositive

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func firstMissingPositive(nums []int) int {

	// Convert negative numbers -> 0
	for i, v := range nums {
		if v < 0 {
			nums[i] = 0
		}
	}

	// Use given array as a hash table by flipping numbers from negative to positive.
	// Indexes of an array represent a number from 1 to the array's length.
	for _, v := range nums {
		vv := abs(v)
		if vv >= 1 && vv <= len(nums) {
			nums[vv-1] = -abs(nums[vv-1])
			if nums[vv-1] == 0 {
				nums[vv-1] = -(len(nums) + 1)
			}
		}
	}

	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] >= 0 {
			return i
		}
	}

	return len(nums) + 1
}

package maximumproductsubarray

/*
	152. Maximum Product Subarray

	Given an integer array nums, find a contiguous non-empty subarray within the array
	that has the largest product, and return the product.
	The test cases are generated so that the answer will fit in a 32-bit integer.
	A subarray is a contiguous subsequence of the array.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	global := nums[0]
	local := 0
	localAbs := 0

	for _, num := range nums {
		if local == 0 || num == 0 {
			local = num
			localAbs = num
			global = max(local, global)
			continue
		}

		if num > 0 {
			local = max(local*num, num)
			localAbs *= num
			global = max(local, global)
			continue
		}

		temp := local * num
		local = max(localAbs*num, num)
		if temp < 0 {
			localAbs = temp
		} else {
			localAbs = num
		}
		global = max(local, global)
	}

	return global
}

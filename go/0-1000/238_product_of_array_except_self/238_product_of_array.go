package productofarrayexceptself

/*
	238. Product of Array Except Self

	Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

	You must write an algorithm that runs in O(n) time and without using the division operation.
*/

func productExceptSelf(nums []int) []int {
	l := len(nums)

	// Create two slices of products from the left and products from the right.
	// Add bounding elements to remove bounds check in the loop.
	leftProd := make([]int, l+1)
	rightProd := make([]int, l+1)
	leftProd[0] = 1
	rightProd[l] = 1

	for i := 0; i < l; i++ {
		leftProd[i+1] = leftProd[i] * nums[i]
		rightProd[l-1-i] = rightProd[l-i] * nums[l-1-i]
	}

	answer := []int{}
	for i := 0; i < l; i++ {
		answer = append(answer, leftProd[i]*rightProd[i+1])
	}

	return answer
}

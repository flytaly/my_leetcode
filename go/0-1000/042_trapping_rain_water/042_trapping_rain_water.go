/*  42. Trapping Rain Water */

package trappingrainwater

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
   1) Move minimal border to the center
   2) Add maximum capacity between borders
   3) Deduct space that was taken by the moving border
*/
func trap(height []int) int {
	l, r := 0, len(height)-1 // border indices
	var result int
	var prevMin, newMin int // store min heights
	var current int         // height of the border that was moved

	for l < r {
		if height[l] == 0 || height[r] == 0 {
			goto skip
		}

		result -= min(prevMin, current) // deduct border height

		newMin = min(height[l], height[r])
		if newMin > prevMin {
			capacity := (r - l - 1) * (newMin - prevMin)
			result += capacity
			prevMin = newMin
		}

	skip:
		if height[l] < height[r] {
			l += 1
			current = height[l]
			continue
		}
		r -= 1
		current = height[r]
	}

	return result
}

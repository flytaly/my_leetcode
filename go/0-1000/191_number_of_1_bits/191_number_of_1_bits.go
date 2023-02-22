package numberof1bits

/*
   Subtracting 1 we change the last 1 to 0 and all 0 on the right end to 1,
   Then using bitwise AND we eliminate all these bits with values 1 on the end.

   Example, 1011:
       1) 1011&1010=>1010
       2) 1010&1001=>1000
       3) 1000&0111=>0000
       => answer is 3
*/
func hammingWeight(num uint32) int {
	count := 0
	for n := num - 1; num > 0; n = num - 1 {
		num = num & n
		count++
	}

	return count
}

// loop with bitshift
func hammingWeight1(num uint32) int {
	count := 0
	for ; num > 0; num = num >> 1 {
		count += int(num % 2)
	}
	return count
}

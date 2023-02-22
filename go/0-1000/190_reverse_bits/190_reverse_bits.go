package reversebits

func reverseBits(num uint32) (res uint32) {
	for i := 0; i < 32; i++ {
		next := (num >> i) & 1  // get the next digit from the right
		res = (res << 1) | next // add to the end
	}
	return res
}

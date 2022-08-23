package sumoftwointegers

func getSum(a int, b int) int {
	a, b = a^b, a&b<<1
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}

// func main() {
// 	fmt.Println(getSum(2, 3))
// 	fmt.Println(getSum(20, 30))
// }

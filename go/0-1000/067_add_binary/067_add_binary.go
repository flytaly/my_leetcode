package addbinary

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	var r = make([]string, len(a))

	carry := 0
	for i := range a {
		i1 := int(a[len(a)-1-i] - '0')
		i2 := 0
		if i < len(b) {
			i2 = int(b[len(b)-1-i] - '0')
		}

		sum := i1 + i2 + carry

		r[len(a)-1-i] = fmt.Sprintf("%v", sum%2)

		carry = 0
		if sum > 1 {
			carry = 1
		}
	}

	join := strings.Join(r, "")
	if carry > 0 {
		return "1" + join
	}

	return join
}

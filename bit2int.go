package bit2int

import (
	"strconv"
)

func BToI(b int) int {
	str := strconv.Itoa(b)
	bs := []byte(str)
	val := 0

	for j, i := uint(0), len(bs)-1; i >= 0; i-- {
		switch string(bs[i]) {
		case "0":
		case "1":
			val += (1 << j)
		default:
			panic("param is err")
		}
		j++
	}

	return val
}

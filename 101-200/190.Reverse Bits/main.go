package main

import (
	"fmt"
	"strconv"
)

func reverseBits(n uint32) uint32 {
	res := uint32(0)
	power := uint32(31)
	for n != 0 {
		res += (n & 1) << power
		n >>= 1
		power--
	}
	return res
}

func main() {
	n := uint32(123)

	fmt.Println(strconv.FormatUint(uint64(n), 2))
	fmt.Println(strconv.FormatUint(uint64(reverseBits(n)), 2))
}

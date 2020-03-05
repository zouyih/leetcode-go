package main

import (
	"bytes"
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var buf bytes.Buffer

	n := len(strs)
	if n == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < n; j++ {
			str := strs[j]
			if i >= len(str) || str[i] != strs[0][i] {
				return buf.String()
			}
		}

		buf.WriteByte(strs[0][i])
	}

	return buf.String()
}

func main() {
	strs := []string{"flower", "flow", "flow"}
	fmt.Println(longestCommonPrefix(strs))
}

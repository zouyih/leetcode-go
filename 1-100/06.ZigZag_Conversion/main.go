package main

import "fmt"

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	res := make([]uint8, 0)
	var ove, even, times int
	var index int
	for i := 0; i < numRows; i++ {
		ove = 2*numRows - 2*i - 2
		even = 2 * i
		index = i
		times = 1
		for index < n {
			if times%2 == 1 && ove != 0 {
				res = append(res, s[index])
			} else if times%2 == 0 && even != 0 {
				res = append(res, s[index])
			}

			if times%2 == 1 {
				index += ove
			} else {
				index += even
			}
			times++

		}
	}

	return string(res)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	trans := make([]string, numRows)

	i := 0
	for i < n {
		for row := 0; row < numRows && i < n; row++ {
			trans[row] = trans[row] + string(s[i])
			i++
		}

		for row := numRows - 2; row >= 1 && i < n; row-- {
			trans[row] = trans[row] + string(s[i])
			i++
		}
	}

	var res string
	for row := 0; row < numRows; row++ {
		res = res + trans[row]
	}

	return res
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	goingDown := false

	curRow := 0
	for _, char := range s {
		rows[curRow] += string(char)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	res := ""
	for _, s := range rows {
		res += s
	}
	return res
}
func main() {
	s := "PAYPALISHIRING"
	fmt.Println(s)
	fmt.Println(convert(s, 3))
	fmt.Println(convert1(s, 3))
	fmt.Println(convert2(s, 3))
}

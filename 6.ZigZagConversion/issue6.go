package __ZigZagConversion

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/25 10:22
   @File: issue6.go
*/

/*
Question Description:
Z字形变换：从上向下，从左往右一次按照Z字形输出指定字符串
*/
func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= 1 {
		return s
	}

	strSlice := make([][]byte, numRows)
	for i := range strSlice {
		strSlice[i] = make([]byte, 1024)
	}

	x, y := 0, 0 //横纵坐标
	forward := true

	for i := range s {
		//从上往下
		if forward {
			strSlice[y][x] = s[i]
			y++
			if y == numRows {
				y--
				forward = false
				continue
			}
		} else {
			y--
			x++
			strSlice[y][x] = s[i]
			if y == 0 {
				y++
				forward = true
				continue
			}
		}
	}
	var result []byte
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(strSlice[i]); j++ {
			if strSlice[i][j] == 0 {
				continue
			}
			result = append(result, strSlice[i][j])
		}
	}
	return string(result)
}

/*leetcode go solution*/
func convert2(s string, numRows int) string {
	slen := len(s)
	if slen < 2 || numRows < 2 {
		return s
	}
	result := make([]byte, slen)

	global_iter := 0

	for i := 0; i < numRows; i++ {
		calc := i
		iter := 0

		if i == numRows-1 {
			iter = 1
		}

		for calc < slen {
			result[global_iter] = s[calc]
			global_iter++

			switch iter {
			case 0:
				if i != 0 {
					iter++
				}
				val := 2 * (numRows - i - 1)
				calc += val
			case 1:
				if i != numRows-1 {
					iter = 0
				}
				calc += 2 * i
			}
		}
	}
	return string(result)
}

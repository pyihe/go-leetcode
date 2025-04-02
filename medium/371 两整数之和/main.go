package main

func GetSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b          // 异或相当于不考虑进位的加法
		carry := (a & b) << 1 // 与运算相当于对应位相加后的进位
		a = sum
		b = carry
	}
	return a
}

func GetSum1(a, b int) int {
	return a + b
}

package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
*/

func main() {
	fmt.Printf("%v\n", climbStairs(10))
}

/*
动态规划:
n = 1 -> 1
n = 2 -> 2
n = 3 -> 3
n = 4 -> 5
.
.
.
可以得出，result(n) = result(n-1)+result(n-2)
*/

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	return cal(n)
}

func cal(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return cal(n-1) + cal(n-2)
}

//分支法
func climbStairs2(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 2 {
		return n
	}
	var result = make([]int, n)
	result[0] = 1
	result[1] = 2
	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n-1]
}

package main

import (
	"fmt"
	"math"
)

/*
整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围[−231, 231− 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321

示例 3：
输入：x = 120
输出：21

示例 4：
输入：x = 0
输出：0

提示：
-231 <= x <= 231 - 1
相关标签：数学
*/

func numReverse(x int) int {
	var res, newRes int
	for x != 0 {
		newRes = res*10 + x%10
		if newRes/10 != res {
			fmt.Println(newRes, res)
			return 0
		}
		res = newRes
		x = x / 10
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}

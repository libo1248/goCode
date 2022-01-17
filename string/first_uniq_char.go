package main

/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

示例：
s = "leetcode"
返回 0

s = "loveleetcode"
返回 2

提示：你可以假定该字符串只包含小写字母。

相关标签：队列 哈希表 字符串 计数
*/

func firstUniqChar(s string) int {
	var cSlice [26]int
	for _, c := range s {
		cSlice[c-'a'] += 1
	}

	for i, c := range s {
		if cSlice[c-'a'] == 1 {
			return i
		}
	}

	return -1
}

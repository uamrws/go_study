package main

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度

func lengthOfLongestSubstring(ss string) int {
	var j int
	mm := make(map[string]int)
	var maxlen int
	for idx, s := range ss {
		v, ok := mm[string(s)]
		if ok && v >= j {
			j = v + 1
		}
		if m := idx + 1 - j; m > maxlen {
			maxlen = m
		}
		mm[string(s)] = idx
	}
	return maxlen
}
func main() {
	lengthOfLongestSubstring("abcabcbb")
}

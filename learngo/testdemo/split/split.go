package split

import "strings"

//Split 用来sep切割字符串s
func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	// 方法一
	n := strings.Count(s, sep) + 1
	ans := make([]string, 0, n)
	n--
	i := 0
	for ; i < n; i++ {
		idx := strings.Index(s, sep)
		ans = append(ans, s[:idx])
		s = s[idx+len(sep):]

	}
	return append(ans, s)

	// 方法二
	//n := strings.Count(s, sep) + 1
	//ans := make([]string, 0, n)
	//sepl := len(sep)
	//start := 0
	//for idx, _ := range s {
	//	if idx+sepl <= len(s) && s[idx:idx+sepl] == sep {
	//		ans = append(ans, s[start:idx])
	//		start = idx + sepl
	//	}
	//}
	//return append(ans, s[start:])
}

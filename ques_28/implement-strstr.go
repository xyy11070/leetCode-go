package ques_28

/**
"aaaaa"
"bba"
**/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hlen, nlen := len(haystack), len(needle)

	for i := 0; i < hlen && i+nlen <= hlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

/** 반복적인 글자 없이 가장 긴 길이의 글자들을 반환 */
func lengthOfLongestSubstring(s string) int {
	result := -1
	words := make(map[byte]int)
	count := -1

	for i := 0; i < len(s); i++ {
		if v, ok := words[s[i]]; !ok {
			words[s[i]] = i
			count++
		} else {
			count = min(i-v-1, count+1)
			words[s[i]] = i
		}
		result = max(result, count)
	}
	return result + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
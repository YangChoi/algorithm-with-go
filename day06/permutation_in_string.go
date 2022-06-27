/** s1이 s2에 포함되는 경우 true, 그렇지 않은 경우 false 반환 */
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var pattern, arr [26]int
	for i := 0; i < len(s1); i++ {
		pattern[int(s1[i]-'a')]++
		arr[int(s2[i]-'a')]++
	}

	i := 0
	for {
		if arr == pattern {
			return true
		}
		if i+len(s1) >= len(s2) {
			break
		}
		arr[int(s2[i]-'a')]--
		arr[int(s2[i+len(s1)]-'a')]++
		i++
	}
	return false
}
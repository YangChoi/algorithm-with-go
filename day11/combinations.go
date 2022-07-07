/**
주워진 두 정수 n과 k,
1부터 n까지의 범위 내에서 k개의 숫자 조합을 반환하라.
*/
func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	var comb func(start int, curComb []int)
	comb = func(start int, curComb []int) {
		if len(curComb) == k {
			// copy of current combination
			dst := make([]int, k)
			copy(dst, curComb)
			result = append(result, dst)
			return
		}

		for i := start; i <= n; i++ {
			curComb = append(curComb, i)
			comb(i+1, curComb)
			curComb = curComb[:len(curComb)-1]
		}
		return
	}
	comb(1, make([]int, 0))
	return result
}
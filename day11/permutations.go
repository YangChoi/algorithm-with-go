/**
정수값이 있는 nums 배열이 주어진다.
가능한 순열을 반환하라.
어떤 순열이던 상관없다.
*/
func permute(nums []int) [][]int {
	// 1 or 0 items
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	// stores final answer
	answer := [][]int{}

	for i := 0; i < len(nums); i++ {
		v := make([]int, len(nums))
		copy(v, nums)

		// swap value i with 0
		// to make removal easier
		v[i], v[0] = v[0], v[i]

		// omit the first value
		v = v[1:]

		// call permute on v
		result := permute(v)

		// loop through result
		for _, j := range result {
			j = append([]int{nums[i]}, j...)
			answer = append(answer, j)
		}

	}
	return answer
}
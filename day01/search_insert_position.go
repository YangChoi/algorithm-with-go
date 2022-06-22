// 오름차순 배열에 target이 들어간 index 구하기
func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums)
	pivot := 0

	for left < right {
		pivot = left + (right-left)/2
		if n := nums[pivot]; n == target {
			return pivot
		} else if n > target {
			right = pivot
		} else {
			left = pivot + 1
		}
	}

	if n := nums[pivot]; n > target {
		return pivot
	} else {
		return pivot + 1
	}
}
// nums 에서 target을 찾는다.
func search(nums []int, target int) int {
	// := 변수 선언 및 대입
	left := 0
	// len() 배열 길이
	right := len(nums) - 1

	for left <= right {
		pivot := left + (right-left)/2
		// left가 right 를 앞지르지 않을 동안
		// pivot은 중간값으로서 작동

		if nums[pivot] > target {
			// 해당 pivot index를 가진 요소가 target보다 크다면
			right = pivot - 1
			// 오른쪽이 pivot 값에서 -1 이 된 값이 된다.
		} else if nums[pivot] < target {
			// 해당 pivot index를 가진 요소가 target 보다 작으면
			left = pivot + 1
			// 왼쪽이 pivot 값에서 +1 이 된 값이 된다.
			// 이 과정을 left가 right보다 작거나 같을 동안 반복한다
		} else {
			return pivot
			// 해당 pivot index 값이 taget값과 같다면 pivot은 우리가 찾던 그 숫자의 index이므로 return
		}
	}

	return -1
	// taget 값이 없는 경우엔 -1 을 반환

}
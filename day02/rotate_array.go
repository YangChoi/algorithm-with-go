/**
주어진 수 만큼 오른쪽으로 순회하기
*/

/**
<배열 선언>
arr1 := [5]int{1, 2, 3, 4, 5}
: 5개라고 명시적으로 크기 지정
arr2 := [...]int(6, 7, 8, 9, 10)
: [...] > 요소갯수만큼 크기 지정

> 크기 변경을 위해 append 사용?
arr1 = append(arr1, 6)
= error

<슬라이스 선언>
slice := []int{1, 2, 3, 4, 5}
: 크기 지정 하지 않음, 가변 배열

> 크기 변경을 위해 append 사용?
slice1 = append(slice1, 6)
= [1 2 3 4 5 6]
*/
func rotate(nums []int, k int) {
	len := len(nums)
	k = k % len
	// 배열 길이 만큼 tmp 슬라이스 생성
	tmp := make([]int, len)

	for j := 0; j < len; j++ {
		if j < len-k {
			tmp[j+k] = nums[j]
		} else {
			tmp[j-len+k] = nums[j]
		}
	}
	copy(nums, tmp)
}
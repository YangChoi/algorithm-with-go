/**
받은 배열의 원소 중 0을 모두 배열의 끝으로 몰아넣기
*/
func moveZeroes(nums []int) {
	length := len(nums)
	index := 0
	count := 0

	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			// 원소가 0이 아닌 경우
			nums[index] = nums[i]
			// 해당 원소는 제자리에 머뭄 (i, index 모두 0부터 시작하므로)
			index++
			// 다음 인덱스로
		} else {
			// 원소가 0일 경우엔
			count++
			// count + 1
		}
	}

	for i := 0; i < count; i++ {
		// for 문은 위에서 받은 count 수 만큼 돈다
		nums[index+i] = 0
		// 배열의 길이 만큼 돈 index에 +i 를 하면 배열의 가장 끝 자리가 된다.
		// 그 자리를 count 수만큼 0으로 채운다.
	}
}
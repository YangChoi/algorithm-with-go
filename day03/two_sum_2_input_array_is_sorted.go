/**
주어진 배열의 원소를 더해 target 이 되도록 만든다.
그리고 그 target 값을 만든 원소들을 찾는다.
*/
func twoSum(numbers []int, target int) []int {
	// 여러 변수에 한 번에 같이 선언 및 초기화
	// 변수 선언과 초기화를 함께 하면 같은 타입의 변수가 아니여도 됨
	// 선언만 하려면 같은 타입만!
	i, j := 0, len(numbers)-1
	for {
		// i는 0부터, j는 배열의 마지막 index 부터
		if numbers[i]+numbers[j] > target {
			// 둘을 더했을 때 target 보다 크다면
			// 큰 값일 j의 index를 왼쪽으로 이동시킴
			j--
		} else if numbers[i]+numbers[j] < target {
			// 둘을 더했을 때 target보다 작다면
			i++
			// 작은 값일 i의 index를 오른쪽으로 이동시킴
		} else {

			return []int{i + 1, j + 1}
		}
	}
	return []int{i + 1, j + 1}
}
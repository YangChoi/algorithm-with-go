/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// 첫번째로 좋지 않았던 버전을 찾는다.

func firstBadVersion(n int) int {
	left := int32(0)
	right := int32(n)
	pivot := int32(0)
	check := false
	answer := int(0)

	for left <= right {
		pivot = left + (right-left)/2
		// & 참조를 위한 해당 변수의 주소값을 매개변수로 전달 ( call by reference )
		isBadVersionChecker(&pivot, &left, &right, &check, &answer)
		if check {
			return answer
		}
	}
	return answer
}

func isBadVersionChecker(pivot, left, right *int32, check *bool, answer *int) {
	pivotLeft := isBadVersion(int(*pivot - 1))
	pivotCenter := isBadVersion(int(*pivot))
	pivotRight := isBadVersion(int(*pivot + 1))

	if pivotCenter {
		if pivotLeft {
			// right는 포인터로 int32 타입을 지정당함 (매개변수의 메모리 주소 전달 : 값 변경 가능)
			*right = *pivot - 1
		} else {
			*check = true
			*answer = int(*pivot)
		}
	} else {
		if !pivotRight {
			*left = *pivot + 1
		} else {
			*check = true
			*answer = int(*pivot + 1)
		}
	}
}

/**
<참고>
- call by value (값에 의한 호출)
값에 의한 호출은 함수를 호출할 때 매개변수 값을 복사해서 함수 내부로 전달
함수 내부에서는 전달된 매개변수의 본래 값을 변경할 수 없음
함수 내부에서 본래값 변경하려면 & 연산자로 변수의 메모리 주소 전달해야함

- call by reference(참조에 의한 호출)
참조에 의한 호출로 매개변수 전달 시 함수에서는 전달한 매개변수의 메모리 주소를 매개변수로 받음
이때 *연산자를 사용해 매개변수 타입을 포인터로 지정해야함
함수의 매개변수타입을 포인터로 지정하면 변수의 값이 아닌 변수의 메모리 주소가 전달됨
이렇게 포인터로 메모리 주소에 접근 시 매개변수로 전달된 인수의 본래 값 변경 가능
단, 참조타입인 슬라이스와 맵은 메모리 참조값을 전달하는 것이 기본
*/

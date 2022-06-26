/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/** linked list의 중간 노드를 찾아서 반환.
만약 중간 노드가 두개가 있다면 두번째 것을 반환
*/
func middleNode(head *ListNode) *ListNode {
	tmp := [100]*ListNode{}
	// ListNode 타입의 슬라이스 만듦
	count := 0

	// head가 nil이 아닐 동안 반복
	for head != nil {
		tmp[count] = head
		// count는 index로 기능하고 그 요소가 head가 됨
		head = head.Next
		// head는 head의 다음 값이 되고
		// 이때 head.Next가 nil이 되는 순간 for문에서 빠져나옴
		// count 역시 그대로 for문을 빠져나옴
		count++
		// count는 ++
	}
	return tmp[count/2]
	// for문을 돌며 head(ListNode) 값을 하나하나 센 count를 2로 나누어 중간값 도출
}
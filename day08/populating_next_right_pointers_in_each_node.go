/**
같은 레벨을 가진 노드 두개가 있다.
모든 부모노드는 2개의 자식 노드를 가진다.
각 노드는 자신의 오른쪽 다음 노드에 포인터를 가지고 있다.
만약 오른쪽 노드가 없다면 해당 포인터는 null이 된다.
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var pre, cur *Node
	pre = root

	for pre.Left != nil {
		cur = pre
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Left

	}
	return root
}
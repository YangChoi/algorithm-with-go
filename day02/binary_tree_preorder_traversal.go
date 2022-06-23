
/**
  자료구조 이진탐색트리 (preorder - 전위순회)
  tree의 node를 지정된 순서대로 방문하는 것
  > 노드 방문 > 왼쪽 서브트리 전위순회 > 오른쪽 서브트리 전위순회
  노드를 먼저 방문하고 왼쪽 끝까지 내려간 다음 오른쪽으로 이동하여 다시 시작하거나
  오른쪽으로 이동해 순회 반복
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// var foo int = 10 (할당만)
// foo := (정의, 할당)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	p := root
	// make 길이가 0인 TreeNode를 만든다
	// TreeNode 메모리 주소에 직접 접근
	stack := make([]*TreeNode, 0)
	// append : 슬라이스 반환 (배열은 사용불가)
	stack = append(stack, p)
	// stack에 위에서 만든 슬라이스와, root 할당

	for len(stack) != 0 {
		// stack 길이가 0이 아닌 동안 반복
		visit := stack[len(stack)-1]
		// 방문한 곳 확인
		stack = stack[:len(stack)-1]
		// stack 할당
		ret = append(ret, visit.Val)
		// 방문한 value 할당
		if visit.Right != nil {
			// 오른쪽 노드가 없다면
			stack = append(stack, visit.Right)
			// stack에 오른쪽 노드 append
		}
		if visit.Left != nil {
			// 왼쪽 노드가 없다면
			stack = append(stack, visit.Left)
			// stack에 왼쪽 노드 append
		}
	}
	return ret
}


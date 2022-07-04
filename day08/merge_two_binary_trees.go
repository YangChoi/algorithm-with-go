/**
노드 두개를 합친다.
겹치는 노드끼리는 sum을 하여 새로운 노드를 만든다.
한 쪽이 null일 경우엔, null 이 아닌 다른 한 쪽 노드로 노드를 만든다.
둘다 null일 경우엔 새로운 노드 트리에 적용시키지 않는다.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	} else {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}
	return root1
}
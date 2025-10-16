package binarytreegeneral

func Flatten(root *TreeNode) {
	flatten(root)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left != nil {
		graftRight(root.Left, root.Right)
		root.Right = root.Left
		root.Left = nil
	}
}

func graftRight(root, branch *TreeNode) {
	if root == nil {
		return
	}

	if root.Right != nil {
		graftRight(root.Right, branch)
		return
	}

	root.Right = branch
}

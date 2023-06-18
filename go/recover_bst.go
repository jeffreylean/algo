package algo

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// swap the wrong
// o(1) space
func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode
	travel(&root, &first, &prev, &second)
	first.Val, second.Val = second.Val, first.Val
}

func travel(root, first, prev, second **TreeNode) {
	if *root == nil {
		return
	}

	travel(&(*root).Left, first, prev, second)
	if *prev != nil {
		if *first == nil && (*prev).Val >= (*root).Val {
			*first = *prev
		}
		if *first != nil && (*prev).Val >= (*root).Val {
			*second = *root
		}
	}
	*prev = *root
	travel(&(*root).Right, first, prev, second)
}

// using sort and reconstruc the who tree
// o(n) space
func sortRecoverTree(root *TreeNode) {
	arr := []int{}
	tree := []*TreeNode{}
	sortTravel(root, &arr, &tree)

	sorted := arr
	sort.Ints(sorted)

	for i, each := range sorted {
		tree[i].Val = each
	}
}

func sortTravel(root *TreeNode, arr *[]int, tree *[]*TreeNode) {
	if root == nil {
		return
	}
	sortTravel(root.Left, arr, tree)
	*arr = append(*arr, root.Val)
	*tree = append(*tree, root)
	sortTravel(root.Right, arr, tree)
}

package _18

func sortedArrayToBST(nums []int) *TreeNode {
	var buildBST func(int, int) *TreeNode
	buildBST = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := left + (right-left)/2
		root := &TreeNode{Val: nums[mid]}

		root.Left = buildBST(left, mid-1)
		root.Right = buildBST(mid+1, right)

		return root
	}

	return buildBST(0, len(nums)-1)
}

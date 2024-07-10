package merge_88

// merge
// @Description 合并两个有序数组,时间复杂度O(m+n),空间复杂度O(1)
// @Author Oberl-Fitzgerald 2024-07-10 18:25:01
// @Param  nums1 []int
// @Param  m int
// @Param  nums2 []int
// @Param  n int
func merge(nums1 []int, m int, nums2 []int, n int) {
	// p1为指向 nums1 中最后一个元素
	p1 := m - 1
	//  p为指向合并后的数组的最后一个元素
	p := m + n - 1
	// p2为指向 nums2 中最后一个元素
	p2 := n - 1
	for p2 >= 0 {
		//  如果 p1 >= 0 并且 nums1[p1] >= nums2[p2]，则将 nums1[p1] 放入 nums1[p] 中，并将 p1 和 p 减 1,否则将 nums2[p2] 放入 nums1[p] 中，并将 p2 和 p 减 1
		//  从后往前遍历，将较大的数放到 nums1 的后面
		if p1 >= 0 && nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
}

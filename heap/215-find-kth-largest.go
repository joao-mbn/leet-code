// https://leetcode.com/problems/kth-largest-element-in-an-array/?envType=study-plan-v2&envId=top-interview-150

package heap

func FindKthLargest(nums []int, k int) int {
	return findKthLargest(nums, k)
}

func findKthLargest(nums []int, k int) int {
	heap := Heapify(nums[:k])

	for _, num := range nums[k:] {
		if num > heap.First() {
			heap.Replace(num)
		}
	}

	return heap.First()
}

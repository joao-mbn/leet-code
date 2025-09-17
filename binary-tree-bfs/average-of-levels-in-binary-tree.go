// https://leetcode.com/problems/average-of-levels-in-binary-tree/?envType=study-plan-v2&envId=top-interview-150

package binarytreebfs

import "leet-code/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *tree.TreeNode) []float64 {
	track := avgTracker(root)
	avgs := []float64{}

	for _, lvSlice := range track {
		total := float64(0)
		if len(lvSlice) == 0 {
			avgs = append(avgs, total)
		}
		for _, value := range lvSlice {
			total += value
		}
		avg := total / float64(len(lvSlice))
		avgs = append(avgs, avg)
	}

	return avgs
}

func avgTracker(root *tree.TreeNode) [][]float64 {
	leftAvg := [][]float64{}
	if root.Left != nil {
		leftAvg = avgTracker(root.Left)
	}

	rightAvg := [][]float64{}
	if root.Right != nil {
		rightAvg = avgTracker(root.Right)
	}

	sizeLeft := len(leftAvg)
	sizeRight := len(rightAvg)
	maxSize := max(sizeLeft, sizeRight)

	lvSlices := [][]float64{{float64(root.Val)}}

	for i := range maxSize {
		lvSlice := []float64{}
		if i < sizeLeft {
			lvSlice = append(lvSlice, leftAvg[i]...)
		}
		if i < sizeRight {
			lvSlice = append(lvSlice, rightAvg[i]...)
		}
		lvSlices = append(lvSlices, lvSlice)
	}

	return lvSlices
}

/* Better algorithm
func averageOfLevels(root *TreeNode) []float64 {
	avgLevels := []float64{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelLength := len(queue)
		avg := float64(0)

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			avg += float64(node.Val)
		}

		avgLevels = append(avgLevels, avg/float64(levelLength))
	}

	return avgLevels
}
*/

package main

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//算法的时间复杂度应该为 O(log (m+n)) 。

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	totalLen := n1 + n2
	a, b := 0, 0
	for a+b < totalLen/2 {
		if a >= n1 {
			b++
			continue
		} else if b >= n2 {
			a++
			continue
		}
		if nums1[a] > nums2[b] {
			b++
		} else {
			a++
		}
	}
	fmt.Println(a, b)
	var max_left, min_right float64
	a_left, b_left := nums1[:a], nums2[:b]
	if len(a_left) == 0 && len(b_left) == 0 {

	} else if len(a_left) == 0 {
		max_left = float64(b_left[len(b_left)-1])
	} else if len(b_left) == 0 {
		max_left = float64(a_left[len(a_left)-1])
	} else {
		max_left = math.Max(
			float64(a_left[len(a_left)-1]),
			float64(b_left[len(b_left)-1]),
		)
	}
	a_right, b_right := nums1[a:], nums2[b:]
	if len(a_right) == 0 && len(b_right) == 0 {

	} else if len(a_right) == 0 {
		min_right = float64(b_right[0])
	} else if len(b_right) == 0 {
		min_right = float64(a_right[0])
	} else {
		min_right = math.Min(
			float64(a_right[0]),
			float64(b_right[0]),
		)
	}

	if totalLen%2 == 1 {
		return min_right
	} else {
		return (max_left + min_right) / 2.0
	}

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
}

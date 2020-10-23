package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sLen := len(nums1) + len(nums2)
	index1, index2, index ,last := 0, 0, 0, 0
	for {
		this := 0
		switch {
		case index1 == len(nums1) :
			this= nums2[index2]
			index2++
		case index2 == len(nums2):
			this= nums1[index1]
			index1++
		case nums1[index1] <= nums2[index2]:
			this = nums1[index1]
			index1++
		default:
			this= nums2[index2]
			index2++
		}
		if index == sLen/2 {
			if sLen %2 == 0{
				return float64(last+this)/2
			}
			return float64(this)
		}
		last = this
		index++
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

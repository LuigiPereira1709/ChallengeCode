package array

import "fmt"

// It's not a very fancy solution, but it's 100% functional and easy to understand
// Don't overcomplicated this solution
func MediaTwoSortedArray() {
	nums1, nums2 := []int{1, 3}, []int{2}

	n1, n2 := len(nums1), len(nums2)
	merged := make([]int, 0, n1+n2)

	i, j := 0, 0

	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	for i < n1 {
		merged = append(merged, nums1[i])
		i++
	}

	for j < n2 {
		merged = append(merged, nums2[j])
		j++
	}

	med, mid := 0.0, len(merged)/2
	if len(merged)%2 == 0 {
		med = (float64(merged[mid-1]) + float64(merged[mid])) / 2.0
	} else {
		med = float64(merged[mid])
	}

	fmt.Println("The median is:", med)
}

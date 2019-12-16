package median

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArraysBruteForce(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	var isOdd bool
	maxToCheck := 0
	median1, median2 := 0, 0
	if total%2 != 0 {
		isOdd = true
		maxToCheck = total / 2
	} else {
		isOdd = false
		maxToCheck = total/2 + 1
	}
	for i, j, k := 0, 0, 0; i <= maxToCheck; i++ {
		median1 = median2
		if nums1[j] < nums2[k] {
			if median2 < nums1[j] {
				median2 = nums1[j]
			}
			if len(nums1)-1 > j {
				j++
			}
		} else {
			if median2 < nums2[k] {
				median2 = nums2[k]
			}
			if len(nums2)-1 > k {
				k++
			}
		}
	}
	if isOdd {
		return float64(median2)
	}
	return float64(median1+median2) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n { // to ensure m<=n
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1 // i is too small
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1 // i is too big
		} else { // i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				if nums2[j] > nums1[i] {
					minRight = nums1[i]
				} else {
					minRight = nums2[j]
				}
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

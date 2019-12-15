package twosum

import (
	"sort"
)

//https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	valueVsIndex := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		theOther := target - nums[i]
		if theOtherIndex, ok := valueVsIndex[theOther]; ok {
			result := []int{i, theOtherIndex}
			sort.Ints(result)
			return result
		}
		valueVsIndex[nums[i]] = i
	}
	return nil
}

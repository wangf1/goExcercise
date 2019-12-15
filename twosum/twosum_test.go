package twosum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testData := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 5, 8, 7, 0, 2, 0, 0, 3, 2, 9, 7, 1},
		[]int{5, 8, 0, 3, 1},
	}

	targets := []int{3, 4, 17, 3}

	results := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{2, 10},
		[]int{2, 3},
	}

	fmt.Printf("-------Test for https://leetcode-cn.com/problems/two-sum----------\n")
	for i := 0; i < len(targets); i++ {
		result := twoSum(testData[i], targets[i])
		fmt.Printf("nums = %v target = %v result = %v\n", testData[i], targets[i], result)
		if result[0] != results[i][0] || result[1] != results[i][1] {
			t.Fatalf("case %d failed. Expected %v but %v\n", i, results[i], result)
		}
	}
}

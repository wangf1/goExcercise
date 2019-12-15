package longestsubstr

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(str string) int {
	charVsIndex := make(map[rune]int)
	startIndex := 0
	maxSubLength := 0
	for i, s := range str {
		dupIndex, ok := charVsIndex[s]
		if ok && dupIndex >= startIndex {
			startIndex = dupIndex + 1
		}
		newMaxSubLength := i - startIndex + 1
		if maxSubLength < newMaxSubLength {
			maxSubLength = newMaxSubLength
		}
		charVsIndex[s] = i
	}
	return maxSubLength
}

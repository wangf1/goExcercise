# https://leetcode-cn.com/problems/longest-palindromic-substring/submissions/


class Solution:

    @staticmethod
    def longestPalindrome(s: str) -> str:
        if not s:
            return ""

        def expand_around_center(left: int, right: int):
            while left >= 0 and right < whole and s[left] == s[right]:
                left -= 1
                right += 1
            return right - left - 1

        start, end, whole = 0, 0, len(s)
        for i in range(whole):
            len1 = expand_around_center(i, i)
            len2 = expand_around_center(i, i + 1)
            length = max(len1, len2)
            if length > end - start:
                start = i - (length - 1) // 2
                end = i + length // 2
        return s[start:end + 1]


if __name__ == '__main__':
    solution = Solution()
    print(solution.longestPalindrome("babad"))
    print(solution.longestPalindrome("bb"))

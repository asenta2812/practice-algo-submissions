class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        window = {}
        count, left = 0, 0

        for right in range(len(s)):
          if s[right] in window:
            left = max(window[s[right]] + 1, left)


          window[s[right]] = right
          count = max(count, right - left + 1)
        return count
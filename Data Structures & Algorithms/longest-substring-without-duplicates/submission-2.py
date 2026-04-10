class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        window = {}
        count, left, right = 0, 0, 0

        while right < len(s):
          if s[right] in window and window[s[right]] >= left:
            left = window[s[right]] + 1


          window[s[right]] = right
          count = max(count, right - left + 1)
          right+=1
        
        return count
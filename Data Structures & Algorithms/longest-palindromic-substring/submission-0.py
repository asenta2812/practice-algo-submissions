class Solution:
    def longestPalindrome(self, s: str) -> str:
      # in ith, check left and right are equal
      n = len(s)
      if n < 2:
        return s

      max_pa_at = -1
      max_pa_len = 0

      def get_max_pa(l: int, r: int):
        nonlocal max_pa_at, max_pa_len
        while l >= 0 and r < n and s[l] == s[r]:
          crr_len = r - l + 1
          if crr_len > max_pa_len:
            max_pa_len = crr_len
            max_pa_at = l
          l-= 1
          r+= 1

      for i in range(n):
        # ababd -> odd
        get_max_pa(i, i)
        
        # abbac -> even
        get_max_pa(i, i + 1)
      

      return s[max_pa_at: max_pa_at + max_pa_len]


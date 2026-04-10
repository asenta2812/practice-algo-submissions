class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        count = max_req = l = 0
        max_reqs = defaultdict(int)

        for r in range(len(s)):
          len_window = r - l + 1
          max_reqs[s[r]] += 1
          max_req = max(max_reqs[s[r]], max_req)
          
          # len - max_req <= k
          if len_window - max_req <= k:
            count = max(count, len_window)
          else:
            max_reqs[s[l]]-=1
            l+=1
        return count
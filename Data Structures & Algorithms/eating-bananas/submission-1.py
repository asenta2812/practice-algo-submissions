class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        maxpile = piles[0]
        for pile in piles[1:]:
          if pile > maxpile:
            maxpile = pile
        
        lower, upper = 1, maxpile
        while lower < upper:
          k = (lower + upper) // 2

          totaltime = 0
          for pile in piles:
            totaltime += (pile + k - 1) // k

          if totaltime > h:
            lower = k + 1
          else:
            upper = k
        
        return upper
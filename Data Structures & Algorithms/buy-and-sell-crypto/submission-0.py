class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit, min_buy = 0, 100

        for sell in prices:
          max_profit = max(max_profit, sell - min_buy)
          min_buy = min(min_buy, sell)
        
        return max_profit
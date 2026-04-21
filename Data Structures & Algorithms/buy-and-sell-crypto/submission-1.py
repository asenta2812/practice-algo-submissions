class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit, min_buy = 0, prices[0]

        for sell in prices[1:]:
          max_profit = max(max_profit, sell - min_buy)
          min_buy = min(min_buy, sell)
        
        return max_profit
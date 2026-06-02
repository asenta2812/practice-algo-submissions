class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = []
        n = len(temperatures)
        res = [0] * n

        for ci in range(n):
            while stack and temperatures[ci] > temperatures[stack[-1]]:
                i = stack.pop()
                res[i] = ci - i
            stack.append(ci)
        
        return res
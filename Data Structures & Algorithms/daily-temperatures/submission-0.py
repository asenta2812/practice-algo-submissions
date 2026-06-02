class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = []
        res = [0] * len(temperatures)

        for ci in range(len(temperatures)):
            curr = temperatures[ci]
            while stack and curr > stack[-1][0]:
                _, i = stack.pop()
                res[i] = ci - i
            stack.append((curr, ci))
        
        return res
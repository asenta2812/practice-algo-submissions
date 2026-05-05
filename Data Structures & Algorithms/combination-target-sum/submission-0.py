class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        res = []
        path = []

        def backtracking(start: int, curr_sum: int):
            # goal:
            if curr_sum == target:
                res.append(list(path))
                return

            # constraint:
            if curr_sum > target:
                return

            # choise:
            for i in range(start, len(nums)):
                path.append(nums[i])
                backtracking(i, curr_sum+nums[i])
                path.pop()
        
        backtracking(0, 0)

        return res

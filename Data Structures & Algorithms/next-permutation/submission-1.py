class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        """
        1. Locate the pivot.
            The pivot is the first character that breaks the non-increasing sequence from the right of the string.
            If no pivot is found, the string is already at its last lexicographical sequence, and the result is just the reverse of the string.
        """
        pivot = len(nums) - 2
        while pivot >= 0 and nums[pivot] >= nums[pivot + 1]:
            pivot -= 1

        if pivot == -1:
            nums.reverse()
            return
       
        """
        2. Find the rightmost successor to the pivot.
        """
        rightmost = len(nums) - 1
        while nums[pivot] >= nums[rightmost]:
            rightmost -= 1

        """
        3. Swap the rightmost successor with the pivot to increase the lexicographical order of the suffix.
        """
        nums[pivot], nums[rightmost] = (nums[rightmost], nums[pivot])

        """
        4. Reverse the suffix after the pivot to minimize its permutation.
        """
        nums[pivot + 1:] = reversed(nums[pivot + 1:])

        
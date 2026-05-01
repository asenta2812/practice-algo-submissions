class Solution:
    def countBits(self, n: int) -> List[int]:
        def hammingWeight(a: int) -> int:
            res = 0
            while a:
                a &= (a-1)
                res += 1
            
            return res

        return [hammingWeight(i) for i in range(n + 1)]
        
        
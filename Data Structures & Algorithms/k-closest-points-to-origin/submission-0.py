class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        maxheap = []
        for [x, y] in points:
            tmp = -(x**2 + y**2)
            heapq.heappush(maxheap, [tmp, x, y])
            if len(maxheap) > k:
                heapq.heappop(maxheap)
        
        res = []
        while maxheap:
            _, x, y = heapq.heappop(maxheap)
            res.append([x, y])
        return res
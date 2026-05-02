"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""

class Solution:
    def minMeetingRooms(self, intervals: List[Interval]) -> int:
        if not intervals:
            return 0
        
        intervals.sort(key=lambda x : x.start)
        room_heaps = []
        heapq.heappush(room_heaps, intervals[0].end)

        for interval in intervals[1:]:
            if interval.start >= room_heaps[0]:
                heapq.heappop(room_heaps)
            heapq.heappush(room_heaps, interval.end)
        return len(room_heaps)


        
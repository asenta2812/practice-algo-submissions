# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        slow, fast = head, head.next
        while fast and fast.next:
          slow = slow.next
          fast = fast.next.next
        
        # cut the list to 2 parts: before slow and after slow
        curr, prev = slow.next, None
        slow.next = None

        # revert the part after slow 
        while curr:
          next_node = curr.next 
          curr.next = prev
          prev = curr
          curr = next_node

        first, second = head, prev
        # first 0 -> 1 -> 2 -> 3 
        # second 6 => 5 => 4
        while second:
          tmp1, tmp2 = first.next, second.next 
          first.next = second
          second.next = tmp1
          first, second = tmp1, tmp2





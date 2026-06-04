# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        total = ListNode(0)

        curr1, curr2, curr3, remember = l1, l2, None, 0
        while curr1 or curr2:
            c1 = curr1.val if curr1 else 0
            c2 = curr2.val if curr2 else 0
            t = c1 + c2 + remember
            new_node = ListNode(t % 10)
            remember = 0
            if t > 9:
                remember = 1
            if not total.next:
                curr3 = total.next = new_node
            else:
                curr3.next = new_node
                curr3 = curr3.next
            
            curr1 = curr1.next if curr1 and curr1.next else None
            curr2 = curr2.next if curr2 and curr2.next else None
        
        if remember == 1:
            curr3.next = ListNode(1)

        return total.next
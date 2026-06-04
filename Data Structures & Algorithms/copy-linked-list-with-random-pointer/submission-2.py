"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return head

        old_to_new = {}

        # curr = head
        # while curr: 
        #     old_to_new[curr] = Node(curr.val)
        #     curr = curr.next

        # curr = head
        # while curr:
        #     old_to_new[curr].next = old_to_new.get(curr.next)
        #     old_to_new[curr].random = old_to_new.get(curr.random)
        #     curr = curr.next
        
        # return old_to_new.get(head)

        curr = head
        while curr:
            if curr not in old_to_new:
                old_to_new[curr] = Node(curr.val)
            
            copy = old_to_new.get(curr)

            
            if curr.next:
                if curr.next not in old_to_new:
                    old_to_new[curr.next] = Node(curr.next.val)
                copy.next = old_to_new.get(curr.next)

            if curr.random:
                if curr.random not in old_to_new:
                    old_to_new[curr.random] = Node(curr.random.val)
                copy.random = old_to_new.get(curr.random)
                
            curr = curr.next
        return old_to_new.get(head)




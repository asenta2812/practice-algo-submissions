# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return head

        # 2. Đệ quy xuống nút cuối cùng
        new_head = self.reverseList(head.next)
        
        # 3. Đảo ngược con trỏ
        # head.next là nút đứng sau, ta trỏ nút sau quay lại head
        head.next.next = head 
        # Ngắt kết nối cũ để tránh vòng lặp vô hạn
        head.next = None 
        
        return new_head

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def sortedListToBST(head: ListNode) -> Optional[TreeNode]:
    """Converts a sorted linked list into a height balanced BST"""
    mid, front, back = splitLinkedListInMiddle(head)

    root = None
    if mid:
        root = TreeNode(mid.val)
        if front:
            root.left = sortedListToBST(front)
        if back:
            root.right = sortedListToBST(back)

    return root


def getLengthOfLinkedList(head):
    itr = head
    count = 0

    while itr:
        count += 1
        itr = itr.next

    return count


def splitLinkedListInMiddle(head):
    """Returns the middle node and front and back linked lists"""
    length = getLengthOfLinkedList(head)

    if length <= 1:
        return head, None, None

    front = head
    itr = head
    pos = 0
    while pos < length // 2:
        pos += 1
        prev = itr
        itr = itr.next

    if prev:
        prev.next = None
    mid = itr
    back = itr.next
    mid.next = None

    return mid, front, back

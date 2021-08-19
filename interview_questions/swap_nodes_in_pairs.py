class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def swapPairs(self, head: ListNode) -> ListNode:
    if head is None:
        return head

    head = swapWithNext(head)
    prevNode = head.next
    currentNode = head.next and head.next.next
    while currentNode:
        currentNode = swapWithNext(currentNode)

        if prevNode is not None:
            prevNode.next = currentNode
            prevNode = currentNode.next

        currentNode = currentNode.next and currentNode.next.next

    return head


def swapWithNext(node: ListNode) -> ListNode:
    if node is None:
        return node

    if node.next is None:
        return node

    nextNode = node.next
    nextNextNode = nextNode.next
    nextNode.next = node
    node.next = nextNextNode

    return nextNode

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def deleteDuplicates(self, head: ListNode) -> ListNode:
    if head is None:
        return None

    counts = getMapOfValueCounts(head)

    front = head
    if front.val in counts and counts[front.val] > 1:
        while front:
            if front.val in counts and counts[front.val] > 1:
                front = front.next
            else:
                break

    if front is None:
        return front

    itr = front.next
    prev = front
    while itr:
        if itr.val in counts and counts[itr.val] > 1:
            prev.next = itr.next
            itr = prev.next
        else:
            prev = itr
            itr = prev.next

    return front


def getMapOfValueCounts(head):
    counts = {}
    itr = head

    while itr:
        if itr.val in counts:
            counts[itr.val] += 1
        else:
            counts[itr.val] = 1

        itr = itr.next

    return counts

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
    lists = splitListIntoKSizeGroups(head, k)
    for i in range(len(lists)):
        lists[i] = reverseIfKLen(lists[i], k)
        print(lists[i])

    return joinLinkedLists(lists)


def joinLinkedLists(lists):
    front = lists[0]

    for i in range(len(lists)):
        lastNode = getLastNode(lists[i])
        if i + 1 < len(lists):
            lastNode.next = lists[i + 1]

    return front


def getLastNode(ll):
    itr = ll
    while itr.next:
        itr = itr.next

    return itr


def reverseIfKLen(ll, k):
    length = getLength(ll)
    if length < 2 or length < k:
        return ll

    prev = ll
    current = ll.next
    prev.next = None

    while current:
        nextNode = current.next
        current.next = prev

        prev = current
        current = nextNode

    return prev


def getLength(ll):
    count = 0
    itr = ll
    while itr:
        count += 1
        itr = itr.next

    return count


def splitListIntoKSizeGroups(head, k):
    lists = []
    front = head
    itr = head
    size = 1

    while itr:
        if size == k:
            nextFront = itr.next
            itr.next = None
            lists.append(front)
            front = nextFront
            itr = nextFront
            size = 1
        else:
            itr = itr.next
            size += 1

    if front:
        lists.append(front)

    return lists

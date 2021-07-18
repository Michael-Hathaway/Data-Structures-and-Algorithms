import heapq
from typing import List


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def mergeKLists(self, lists: List[ListNode]) -> ListNode:
    heap = MinHeap()

    for i in range(len(lists)):
        front = lists[i]
        if front:
            lists[i] = lists[i].next
            heap.add((front.val, i, front))

    if heap.isEmpty():
        return None

    val, index, node = heap.pop()
    if lists[index]:
        heap.add((lists[index].val, index, lists[index]))

    newFront = node
    itr = node

    while not heap.isEmpty():
        val, index, newNode = heap.pop()
        if lists[index]:
            heap.add((lists[index].val, index, lists[index]))
            lists[index] = lists[index].next

        itr.next = newNode
        itr = itr.next
        itr.next = None

    return newFront


class MinHeap:
    def __init__(self):
        self.heap = []

    def isEmpty(self):
        return len(self.heap) == 0

    def add(self, valueNodeIndexTuple):
        heapq.heappush(self.heap, valueNodeIndexTuple)

    def pop(self):
        return heapq.heappop(self.heap)

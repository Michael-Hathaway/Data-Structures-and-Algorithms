import heapq
from typing import List


def findKthLargest(nums: List[int], k: int) -> int:
    heap = MaxHeap()
    for num in nums:
        heap.push(num)

    for _ in range(k):
        val = heap.pop()

    return val


class MaxHeap:
    def __init__(self):
        self.heap = []

    def push(self, value):
        heapq.heappush(self.heap, -value)

    def pop(self):
        if len(self.heap):
            value = heapq.heappop(self.heap)
            return -value

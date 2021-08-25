from typing import List


def merge(self, intervals: List[List[int]]) -> List[List[int]]:
    intervals.sort(key=lambda arg: arg[0])
    finalIntervals = []
    i = 0
    while i < (len(intervals)):
        currentInterval = intervals[i]
        if i == len(intervals) - 1:
            finalIntervals.append(currentInterval)
            break

        j = i + 1
        while j < len(intervals) and intervals[j][0] <= currentInterval[1]:
            currentInterval[1] = max(currentInterval[1], intervals[j][1])
            j += 1

        finalIntervals.append(currentInterval)
        i = j

    return finalIntervals

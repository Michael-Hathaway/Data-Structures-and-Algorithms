def maxSumIncreasingSubsequence(array):
    if len(array) == 0:
        return 0, []

    if len(array) == 1:
        return array[0], array

    sums = [float("-inf") for _ in range(len(array))]
    sums[0] = array[0]
    prev = {}
    prev[0] = None
    maxSumIndex = 0

    for i in range(1, len(array)):
        sums[i] = array[i]

        for j in range(0, i + 1):
            if array[j] < array[i] and sums[j] + array[i] >= sums[i]:
                sums[i] = array[i] + sums[j]
                prev[i] = j
            if sums[i] > sums[maxSumIndex]:
                maxSumIndex = i

    return max(sums), buildSequence(array, prev, maxSumIndex)


def buildSequence(array, prev, maxSumIndex):
    sequence = []
    sequence.append(array[maxSumIndex])

    while maxSumIndex in prev and prev[maxSumIndex] is not None:
        idx = prev[maxSumIndex]
        sequence.append(array[idx])

        maxSumIndex = idx

    return list(reversed(sequence))

from typing import List


def numIslands(self, grid: List[List[str]]) -> int:
    count = 0
    visited: set = set()
    for row in range(len(grid)):
        for col in range(len(grid[0])):
            if grid[row][col] == "1":
                count += 1
                convertConnectedSquaresToZero(grid, row, col, visited)

    return count


def convertConnectedSquaresToZero(grid, startRow, startCol, visited):
    queue = [(startRow, startCol)]

    while queue:
        row, col = queue.pop(0)

        if (row, col) in visited:
            continue

        grid[row][col] = "0"
        visited.add((row, col))

        neighborCells = getAdjacentCells(grid, row, col)
        queue.extend(neighborCells)


def getAdjacentCells(grid, row, col):
    leftBoundary, lowerBoundary = 0, 0
    upperBoundary, rightBoundary = len(grid), len(grid[0])

    cells = []

    if row + 1 < upperBoundary and grid[row + 1][col] == "1":
        cells.append((row + 1, col))

    if row - 1 >= leftBoundary and grid[row - 1][col] == "1":
        cells.append((row - 1, col))

    if col + 1 < rightBoundary and grid[row][col + 1] == "1":
        cells.append((row, col + 1))

    if col - 1 >= lowerBoundary and grid[row][col - 1] == "1":
        cells.append((row, col - 1))

    return cells

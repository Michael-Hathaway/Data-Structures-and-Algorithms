from typing import List


def isValidSudoku(board: List[List[str]]) -> bool:
    allValidSubSquares = validateAllSubSquares(board)
    if not allValidSubSquares:
        return False

    for i in range(0, 9):
        if not isRowValid(board, i) or not isColumnValid(board, i):
            return False

    return True


def isRowValid(board, row):
    nums = set()
    for col in range(0, 9):
        value = board[row][col]
        if value == ".":
            continue
        if value in nums:
            return False
        nums.add(value)

    return True


def isColumnValid(board, col):
    nums = set()
    for row in range(0, 9):
        value = board[row][col]
        if value == ".":
            continue
        if value in nums:
            return False
        nums.add(value)

    return True


def isSubSquareValid(board, topLeftRow, topLeftCol):
    nums = {str(i): False for i in range(1, 10)}

    for row in range(topLeftRow, topLeftRow + 3):
        for col in range(topLeftCol, topLeftCol + 3):
            value = board[row][col]
            if value == ".":
                continue

            if nums[value] is True:
                return False

            nums[value] = True

    return True


def validateAllSubSquares(board):
    for topLeftRow in range(0, 9, 3):
        for topLeftCol in range(0, 9, 3):
            if not isSubSquareValid(board, topLeftRow, topLeftCol):
                return False

    return True

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def isValidBST(root: TreeNode) -> bool:
    return isValidHelper(root, float('-inf'), float('inf'))


def isValidHelper(node, lowerBound, upperBound):
    if node is None:
        return True

    if node.val >= upperBound:
        return False

    if node.val <= lowerBound:
        return False

    if node.left and node.left.val > node.val:
        return False

    if node.right and node.right.val < node.val:
        return False

    isLeftValid = isValidHelper(node.left, lowerBound, node.val)
    isRightValid = isValidHelper(node.right, node.val, upperBound)

    return isLeftValid and isRightValid

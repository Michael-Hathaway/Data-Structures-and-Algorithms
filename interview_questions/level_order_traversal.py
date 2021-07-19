from typing import Dict, List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def levelOrder(self, root: TreeNode) -> List[List[int]]:
    levels: Dict = {}
    getMapOfLevelToNodes(root, levels, 0)

    nodes = []
    for i in range(len(levels)):
        thisLevelNodes = levels[i]
        nodes.append(thisLevelNodes)

    return nodes


def getMapOfLevelToNodes(tree, levels, currentLevel):
    if tree:
        getMapOfLevelToNodes(tree.left, levels, currentLevel + 1)

        if currentLevel in levels:
            levels[currentLevel].append(tree.val)
        else:
            levels[currentLevel] = [tree.val]

        getMapOfLevelToNodes(tree.right, levels, currentLevel + 1)

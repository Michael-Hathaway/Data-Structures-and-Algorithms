class BinaryTree:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right


def flattenBinaryTree(root):
    # Write your code here.
    nodesInOrder = []
    getNodesInOrder(root, nodesInOrder)

    first = nodesInOrder[0]
    for i in range(0, len(nodesInOrder)):
        currentNode = nodesInOrder[i]
        if i > 0:
            currentNode.left = nodesInOrder[i-1]
        if i < len(nodesInOrder)-1:
            currentNode.right = nodesInOrder[i+1]

    return first


def getNodesInOrder(tree, nodes):
    if tree:
        getNodesInOrder(tree.left, nodes)
        nodes.append(tree)
        getNodesInOrder(tree.right, nodes)

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def isSubPath(self, head: ListNode, root: TreeNode) -> bool:
    array = convertLinkedListToList(head)
    found = isSubPathHelper(root, array)

    return found


def searchForPath(tree, nodes):
    if tree is None or tree.val != nodes[0]:
        return False

    if tree.val == nodes[0] and len(nodes) == 1:
        return True

    left = searchForPath(tree.left, nodes[1:])
    right = searchForPath(tree.right, nodes[1:])

    return left or right


def isSubPathHelper(tree, nodes):
    if tree:
        if tree.val == nodes[0]:
            foundHere = searchForPath(tree, nodes)
            if foundHere:
                return True

        foundLeft = isSubPathHelper(tree.left, nodes)
        foundRight = isSubPathHelper(tree.right, nodes)
        return foundRight or foundLeft

    return False


def convertLinkedListToList(head):
    itr = head
    array = []

    while itr:
        array.append(itr.val)
        itr = itr.next

    return array

class Node:
    def __init__(self,value):
        self.left = None
        self.right = None
        self.val = value

def inorder(root):
    if root:
        inorder(root.left)
        print(root.val)
        inorder(root.right)

root = Node(1)
root.left = Node(2)
root.right = Node(3)
root.left.left = Node(4)
root.left.right = Node(5)

print(inorder(root))

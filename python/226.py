# 226. Invert Binary Tree

def solution(root):
    if root is None:
        return root
    
    temp = root.left
    root.left = root.right
    root.right = temp
    
    solution(root.left)
    solution(root.right)

    return root
# fast and slow pointer solution to finding the middle node in a linkedlist

class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

def find_middle(head):
    slow_pointer = head
    fast_pointer = head

    while fast_pointer is not None and fast_pointer.next is not None:
        slow_pointer = slow_pointer.next
        fast_pointer = fast_pointer.next.next

    return slow_pointer

# Example usage
head = Node(1)
head.next = Node(2)
head.next.next = Node(3)
head.next.next.next = Node(4)
head.next.next.next.next = Node(5)

middle = find_middle(head)
print("Middle node value:", middle.value)  # Output: Middle node value: 3

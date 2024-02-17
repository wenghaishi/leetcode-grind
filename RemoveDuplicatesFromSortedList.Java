
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.


public class RemoveDuplicatesFromSortedList {
    public static void main(String[] args) {
    }

    public static ListNode solution(ListNode head) {
        ListNode current = head;
        while (current != null && current.next != null) {
            if (current.next.val == current.val) {
                current.next = current.next.next;
            } else {
                current = current.next;
            }
        }
        return head;
    }
}

// explanation:
// We first create a pointer: current, that points to the head of the linkedlist. Then we iterate throught he linkedlist.
// while the current node and next node is not null, we check if the value of the current node equals to the value of the next node.
// If they are equal, we point the next node to the next next node, effectively removing the duplicate node.
// If they are not equal, we move the current pointer to the next node to continue the iteration
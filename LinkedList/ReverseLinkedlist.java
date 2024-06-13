package LinkedList;

public class ReverseLinkedlist {

    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(3);

        node1.next = node2;
        node2.next = node3;

        solution(node1);
    }

    public static ListNode solution(ListNode head) {
        ListNode current = head;
        ListNode prev = null;

        while (current != null) {
            ListNode temp = current.next;
            current.next = prev;
            current = temp;
            prev = current;
        }

        return prev;
    }
}

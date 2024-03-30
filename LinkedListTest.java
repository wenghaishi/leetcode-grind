

public class LinkedListTest {

    public static void main(String[] args) {
        ListNode head = new ListNode(1);
        ListNode current = head;
        for (int i = 2; i < 6; i++) {
            current.next = new ListNode(i);
            current = current.next;
        }
        ListNode result = solution(head);
        while (result != null) {
            System.out.println(result.val);
            result = result.next;
        }
    }

    public static ListNode solution(ListNode head) {
        ListNode temp = null;
        while (head.next != null && head.next.next != null) {
            temp = head;
            head = head.next;
            head.next = temp;
            head = head.next.next;
        }

        return head;
    }
}

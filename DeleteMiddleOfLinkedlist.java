public class DeleteMiddleOfLinkedlist {
    public static void main(String[] args) {

    }

    public static ListNode solution(ListNode head) {
        if (head.next == null) {
            return null;
        }
        if (head.next.next == null) {
            head.next = null;
            return head;
        }

        ListNode slow = head;
        ListNode fast = head;
        while (fast != null && fast.next.next != null) {
            if (fast.next.next.next == null) {
                slow.next = slow.next.next;
                return head;
            }
            slow = slow.next;
            fast = fast.next.next;
            if (fast.next.next == null) {
                slow.next = slow.next.next;
                return head;
            }
        }
        return null;
    }
}

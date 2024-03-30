public class RemoveLinkedlistElements {
    public static void main(String[] args) {

    }

    public static ListNode solution(ListNode head, int val) {
        ListNode current = head;
        while (current != null && current.next != null) {
            if (current.next.val == val) {
                current.next = current.next.next;
            } else {
                current = current.next;
            }
        }
        return head;
    }
}

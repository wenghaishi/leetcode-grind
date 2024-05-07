package LinkedList;

public class RemoveNodeFromList {
    public static void main(String[] args) {

    }

    public ListNode removeNodes(ListNode head) {
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode current = dummy;
        while (current.next != null) {
            ListNode temp = current;
            while (temp.next != null) {
                if (temp.next >= current)
            }

        }

    }
}

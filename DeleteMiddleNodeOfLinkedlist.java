public class DeleteMiddleNodeOfLinkedlist {
  public static void main(String[] args) {

  }

  public static ListNode solution(ListNode head) {
    ListNode slow = head.next, fast = head.next.next;

    while (fast != null && fast.next != null) {
      fast = fast.next.next;
      slow = slow.next;
    }
    slow.next = slow.next.next;
    return head;
  }
}


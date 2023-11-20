public class LinkedListCycle {
  public static void main(String[] args) {

  }

  public static boolean solution(ListNode head) {
    ListNode slow = head;
    ListNode fast = head;
    while(fast.next.next != null && fast != null) {
      fast = fast.next.next;
      slow = slow.next;
      if (slow == fast) {
        return true;
      }
    }
    return false;
  }
}

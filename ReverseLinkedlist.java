// Given the head of a singly linked list, reverse the list, and return the reversed list.

public class ReverseLinkedlist {
  public static void main(String[] args) {
  }

  public static ListNode solution(ListNode head) {
    ListNode newHead = null;
    while (head != null) {
      ListNode next = head.next;
      head.next = newHead;
      newHead = head;
      head = next;
    }
    return newHead;
  }
}

package LinkedList;

import java.util.ArrayList;

public class PalinedromeList {
    public static void main(String[] args) {
        // Test 1: Palindrome with odd length
        ListNode head1 = new ListNode(1);
        head1.next = new ListNode(2);
        head1.next.next = new ListNode(1);
        System.out.println("Test 1 (Odd Length Palindrome): " + isPalindrome(head1));

        // Test 2: Palindrome with even length
        ListNode head2 = new ListNode(1);
        head2.next = new ListNode(2);
        head2.next.next = new ListNode(2);
        head2.next.next.next = new ListNode(1);
        System.out.println("Test 2 (Even Length Palindrome): " + isPalindrome(head2));

        // Test 3: Non-palindrome list
        ListNode head3 = new ListNode(1);
        head3.next = new ListNode(2);
        head3.next.next = new ListNode(3);
        System.out.println("Test 3 (Non-Palindrome): " + isPalindrome(head3));
    }

    public static boolean isPalindrome(ListNode head) {
        ArrayList<Integer> myArr = new ArrayList<Integer>();
        while (head != null) {
            myArr.add(head.val);
            head = head.next;
        }
        int start = 0;
        int end = myArr.size() - 1;
        while (end > start) {
            if (myArr.get(start) != myArr.get(end)) {
                return false;
            } else {
                start++;
                end--;
            }
        }
        return true;
    }
}

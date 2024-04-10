package LinkedList;

import java.util.ArrayList;

public class NextLargerNode {

    public static void main(String[] args) {
        ListNode head1 = new ListNode(1);
        head1.next = new ListNode(2);
        head1.next.next = new ListNode(3);

        int[] head1Result = nextLargerNodes(head1);

        for (int i = 0; i < head1Result.length; i++) {
            System.out.println(head1Result[i]);
        }

    }

    public static int[] nextLargerNodes(ListNode head) {
        ArrayList<Integer> myArr = new ArrayList<Integer>();

        ListNode current = head;
        while (current != null) {
            ListNode pointerA = current.next;
            while (pointerA != null) {
                if (current.val < pointerA.val) {
                    myArr.add(pointerA.val);
                    break;
                }
                pointerA = pointerA.next;
            }
            if (pointerA == null) {
                myArr.add(0);
            }
            current = current.next;
        }

        int[] result = new int[myArr.size()];
        for (int i = 0; i < myArr.size(); i++) {
            result[i] = myArr.get(i);
        }
        return result;
    }
}

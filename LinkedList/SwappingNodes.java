package LinkedList;

import java.util.ArrayList;

public class SwappingNodes {
    public static void main(String[] arg) {

    }

    public static ListNode solution(ListNode head, int k) {
        ArrayList<Integer> myArr = new ArrayList<Integer>();

        ListNode current = head;
        while (current != null) {
            myArr.add(current.val);
            current = current.next;
        }

        int temp = myArr.get(myArr.size() - k);
        myArr.set(myArr.size() - k, myArr.get(k - 1));
        myArr.set(k - 1, temp);

        ListNode dummy = new ListNode(0);

        ListNode pointer = new ListNode(myArr.get(0));
        dummy.next = pointer;
        for (int j = 1; j < myArr.size(); j++) {
            pointer.next = new ListNode(myArr.get(j));
            pointer = pointer.next;
        }

        return dummy.next;

    }
}

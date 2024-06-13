package LinkedList;

import java.util.ArrayList;
import java.util.Collections;

public class MergeTwoSortedLists {

    public static void main(String args[]) {
        // Test case 1: Both lists are non-empty
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(4);
        list1.next.next = new ListNode(5);

        ListNode list2 = new ListNode(1);
        list2.next = new ListNode(3);
        list2.next.next = new ListNode(4);

        ListNode mergedList = solution(list1, list2);
        printList(mergedList); // Expected output: 1 1 3 4 4 5

        // Test case 2: One list is empty
        ListNode list3 = null;

        ListNode list4 = new ListNode(2);
        list4.next = new ListNode(6);

        mergedList = solution(list3, list4);
        printList(mergedList); // Expected output: 2 6

        // Test case 3: Both lists are empty
        ListNode list5 = null;
        ListNode list6 = null;

        mergedList = solution(list5, list6);
        printList(mergedList); // Expected output: (nothing)

        // Test case 4: Lists with negative and positive numbers
        ListNode list7 = new ListNode(-1);
        list7.next = new ListNode(0);
        list7.next.next = new ListNode(1);

        ListNode list8 = new ListNode(-3);
        list8.next = new ListNode(2);
        list8.next.next = new ListNode(3);

        mergedList = solution(list7, list8);
        printList(mergedList); // Expected output: -3 -1 0 1 2 3
    }

    public static ListNode solution(ListNode list1, ListNode list2) {
        ListNode currentList1 = list1;
        ListNode currentList2 = list2;

        ArrayList<Integer> myList = new ArrayList<Integer>();

        while (currentList1 != null) {
            myList.add(Integer.valueOf(currentList1.val));
            currentList1 = currentList1.next;
        }

        while (currentList2 != null) {
            myList.add(Integer.valueOf(currentList2.val));
            currentList1 = currentList2.next;
        }

        Collections.sort(myList);

        ListNode solution = new ListNode(myList.get(0));
        ListNode current = solution;

        for (int i = 1; i < myList.size(); i++) {
            ListNode nextNode = new ListNode(myList.get(i));
            current.next = nextNode;
            current = current.next;
        }

        return solution;
    }

    public static void printList(ListNode node) {
        while (node != null) {
            System.out.print(node.val + " ");
            node = node.next;
        }
        System.out.println();
    }
}

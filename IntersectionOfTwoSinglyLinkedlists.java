import java.util.HashSet;

public class IntersectionOfTwoSinglyLinkedlists {

    static class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
            next = null;
        }
    }

    public static void main(String[] args) {
        ListNode headA = new ListNode(4);
        ListNode headB = new ListNode(5);
        ListNode node1 = new ListNode(8);
        ListNode node2 = new ListNode(1);
        ListNode node3 = new ListNode(4);
        ListNode node4 = new ListNode(5);
        ListNode node5 = new ListNode(6);
        ListNode node6 = new ListNode(1);
        headA.next = node1;
        headA.next.next = node2;
        headA.next.next.next = node3;
        headA.next.next.next.next = node4;
        headA.next.next.next.next.next = node5;
        headA.next.next.next.next.next.next = node6;
        headB.next = node4;
        headB.next.next = node5;
        headB.next.next.next = node6;
        System.out.println(solution(headA, headB).val);
    }

    public static ListNode solution(ListNode headA, ListNode headB) {
        HashSet<ListNode> mySet = new HashSet<>();
        ListNode currentA = headA;
        ListNode currentB = headB;
        while (currentA != null && currentA.next != null) {
            mySet.add(currentA);
            currentA = currentA.next;
        }
        while (currentB != null && currentB.next != null) {
            if (mySet.contains(currentB)) {
                return currentB;
            } else {
                currentB = currentB.next;
            }
        }
        return null;
    }
}

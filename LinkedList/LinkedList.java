package LinkedList;

public class LinkedList {

    static class Node {
        int value;
        Node next;

        public Node(int value) {
            this.value = value;
        }
    }

    Node root;

    public LinkedList(int value) {
        Node newNode = new Node(value);
        newNode.next = null;
        root = newNode;
    }

    public LinkedList() {
    }

    public void addNode(int value) {
        Node newNode = new Node(value);

        Node temp = root;
        while (temp.next != null) {
            temp = temp.next;
        }

        temp.next = newNode;
    }

    public void printList() {
        Node temp = root;
        while (temp != null) {
            System.out.println(temp.value);
            temp = temp.next;
        }
    }

    public static void main(String[] args) {
        LinkedList myList = new LinkedList(5);
        myList.addNode(4);
        myList.addNode(3);
        myList.addNode(2);
        myList.addNode(1);
        myList.printList();
        myList.reverseLinkedList();
        myList.printList();

    }

    public void reverseLinkedList() {
        Node current = this.root;
        Node next = null;
        Node prev = null;
        while (current != null) {
            next = current.next;
            current.next = prev;
            prev = current;
            current = next;
        }

        this.root = prev;

    }

}

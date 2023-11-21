public class StackSort {

    static class Node {
        int value;
        Node next;

        public Node(int value) {
            this.value = value;
        }
    }

    static class Stack {
        private Node top;
        private int length;

        public Stack(int value) {
            Node newNode = new Node(value);
            this.top = newNode;
            this.length = 1;
        }

        public Stack() {

        }

        public void push(int value) {
            Node newNode = new Node(value);
            if (this.length == 0) {
                this.top = newNode;
                this.length = 1;
            } else {
                Node temp = this.top;
                top = newNode;
                top.next = temp;
                length++;
            }
        }

        public void pop() {
            if (this.length == 0) return;
            Node temp = this.top;
            top = top.next;
            temp.next = null;
            length--;
        }
        public void printStack() {
            Node current = this.top;
            while (current != null) {
                System.out.println(current.value);
                current = current.next;
            }
        }

        public boolean isEmpty() {
            return length == 0;
        }

    }

    public static void main(String[] args) {
        Stack myStack = new Stack(1);
        myStack.push(2);
        myStack.push(3);
        // myStack.printStack();
        myStack.pop();
        myStack.printStack();
    }

    public Stack sort(Stack stack) {
        Stack tempStack = new Stack();
        while (!stack.isEmpty()) {
            Node temp = stack.top;
            while (!tempStack.isEmpty()) {
                if ( temp > tempStack.top) {
                     
                }
            }
        }
    }
}

import java.util.Stack;

public class AsteroidCollision {
    public static void main(String[] args) {
        int[] asteroids = { 5, 10, -5 }; // Example input
        int[] result = solution(asteroids);
        for (int num : result) {
            System.out.print(num + " ");
        }
    }

    public static int[] solution(int[] asteroids) {
        Stack<Integer> myStack = new Stack<Integer>();

        for (int i = 0; i < asteroids.length; i++) {
            boolean next = false;

            if (myStack.isEmpty()) {
                myStack.push(asteroids[i]);
            } else {
                while (next != true && !myStack.isEmpty()) {
                    boolean isSameDirection = (asteroids[i] < 0 && myStack.peek() < 0)
                            || (asteroids[i] > 0 && myStack.peek() > 0);
                    if (isSameDirection) {
                        myStack.push(asteroids[i]);
                        next = true;
                    } else {
                        if (asteroids[i] < 0 && asteroids[i] + myStack.peek() < 0) {
                            myStack.pop();
                        } else if (asteroids[i] + myStack.peek() == 0) {
                            myStack.pop();
                            next = true;
                        } else {
                            next = true;
                        }
                    }
                }

            }
        }

        int[] solutionArr = new int[myStack.size()];
        while (!myStack.isEmpty()) {
            for (int i = myStack.size() - 1; i > -1; i--) {
                solutionArr[i] = myStack.pop();
            }
        }

        return solutionArr;
    }
}

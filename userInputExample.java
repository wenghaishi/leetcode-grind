import java.util.Scanner;

public class userInputExample {
    public static void main(String[] args) {
        // Create a Scanner object to read from the console
        Scanner scanner = new Scanner(System.in);

        // Prompt the user for input
        System.out.print("Enter your name: ");

        // Read a line of text as a string
        String name = scanner.nextLine();

        // Prompt the user for another input
        System.out.print("Enter your age: ");

        // Read an integer
        int age = scanner.nextInt();

        // Display the user's input
        System.out.println("Hello, " + name + "! You are " + age + " years old.");

        // Close the scanner to free up resources (optional)
        scanner.close();
    }
}

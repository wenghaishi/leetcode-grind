public class PowerOfTwo {
    public static void main(String[] args) {

    }

    public static boolean solution(int n) {
        return n > 0 && (n & (n - 1)) == 0;
    }
}

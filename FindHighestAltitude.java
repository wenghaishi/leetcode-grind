public class FindHighestAltitude {
    public static void main(String[] args) {
        int[] test = { -5, 1, 5, 0, -7 };

        System.out.println(solution(test));
    }

    public static int solution(int[] gain) {
        int current = 0;
        int highest = 0;
        for (int i = 0; i < gain.length; i++) {
            current += gain[i];
            if (current > highest)
                highest = current;
        }

        return highest;
    }
}
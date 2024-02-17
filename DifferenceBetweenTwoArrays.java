import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

public class DifferenceBetweenTwoArrays {
    public static void main(String[] args) {
    }

    public static List<List<Integer>> solution(int[] nums1, int[] nums2) {
        HashSet<Integer> commonIntegers = new HashSet<>();

        for (int i = 0; i < nums1.length; i++) {
            for (int j = 0; j < nums2.length; j++) {
                if (nums1[i] == nums2[j]) {
                    commonIntegers.add(nums1[i]);
                }
            }
        }

        
        ArrayList<Integer> nums1List = new ArrayList<>();
        for (int i = 0; i < nums1.length; i++) {
            if (!commonIntegers.contains(nums1[i])) {
                nums1List.add(nums1[i]);
            }
        }
        ArrayList<Integer> nums2List = new ArrayList<>();
        for (int i = 0; i < nums2.length; i++) {
            if (!commonIntegers.contains(nums2[i])) {
                nums2List.add(nums2[i]);
            }
        }
        return List.of(nums1List, nums2List);
    }
}

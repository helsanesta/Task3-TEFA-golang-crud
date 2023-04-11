import java.util.ArrayList;
import java.util.Arrays;

class experiment8 {
    public static ArrayList<ArrayList<String>> results = new ArrayList<>();
    public static String[] candidates = { "Baswedan", "Subianto", "Maharani" };

    public static void main(String[] args) {
        ArrayList<ArrayList<String>> chairs = arrange(candidates);
        System.out.println(Arrays.deepToString(chairs.toArray()));
    }

    public static ArrayList<ArrayList<String>> arrange(String[] array) {
        return arrange(array, new ArrayList<>());
    }

    public static ArrayList<ArrayList<String>> arrange(String[] array, ArrayList<String> memory) {
        String current;
        ArrayList<String> newMemory;
        if (memory == null) {
            newMemory = new ArrayList<>();
        } else {
            newMemory = new ArrayList<>(memory);
        }
        for (int i = 0; i < array.length; i++) {
            current = array[i];
            ArrayList<String> newArray = new ArrayList<>(Arrays.asList(array));
            newArray.remove(i);
            if (array.length == 1) {
                newMemory.add(current);
                results.add(new ArrayList<>(newMemory));
                return results;
            }
            newMemory.add(current);
            arrange(newArray.toArray(new String[0]), newMemory);
            newArray.add(i, current);
            newMemory.remove(newMemory.size() - 1);
        }
        return results;
    }
}

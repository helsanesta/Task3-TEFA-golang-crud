class experiment7 {
    public static void main(String[] args) {
        String[] foods = { "Sate", "Bakso", "Dimsum", "Rames" };
        String[] drinks = { "Jeruk", "Teh", "Kelapa", "Cendol" };
        logPairs(foods, drinks, "Menu");
    }

    public static void logPairs(String[] array1, String[] array2, String words) {
        int counter = 0;
        for (int i = 0; i < array1.length; i++) {
            for (int j = 0; j < array2.length; j++) {
                counter++;
                System.out.println(words + ' ' + counter + ' ' + array1[i] + " dan " + array2[j]);
            }
        }
    }
}
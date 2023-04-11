import java.util.Arrays;

class experiment3 {
    public static void main(String[] args) {
        String company = "telkom";
        String[] companies = new String[100000];
        Arrays.fill(companies, company);
        findCompany(companies);
    }

    public static void findCompany(String[] array) {
        long tx = System.nanoTime();
        for (int i = 0; i < array.length; i++) {
            if (array[i] == "telkom") {
                System.out.println("Company Found!");
            }
        }
        long ty = System.nanoTime();
        double duration = (ty - tx) / 1_000_000.0;
        System.out.println("The performance is " + duration + " ms");
    }
}
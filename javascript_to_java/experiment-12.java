import java.util.ArrayList;

class experiment12 {
    public static ArrayList<String> generateData(int n) {
        String baseNumber = "882";
        ArrayList<String> customers = new ArrayList<String>();
        String mobileNumber;

        for (int i = 0; i < n; i++) {
            mobileNumber = baseNumber + String.format("%09d", 1);
            customers.add(mobileNumber);
        }

        return customers;
    }

    public static void sendPromoDiscount(ArrayList<String> input) {
        for (int i = 0; i < input.size(); i++) {
            System.out.println("Sending Promo to " + input.get(i)); // O(n)
        }
        System.out.println("Its Finished to send Promo to " + input.size() + " Customers"); // 0(1)

        for (int i = 0; i < input.size(); i++) {
            System.out.println("Sending Discount to " + input.get(i)); // O(n)
        }
        System.out.println("Its Finished to send Discount to " + input.size() + " Customers"); // 0(1)
    }

    public static void main(String[] args) {
        ArrayList<String> data = generateData(1000);
        sendPromoDiscount(data);
    }
}

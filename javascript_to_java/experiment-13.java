class experiment13 {
    public static String[] generateData(int n) {
        String baseNumber = "082";
        String[] customers = new String[n];
        String mobileNumber;
        for (int i = 0; i < n; i++) {
            mobileNumber = baseNumber + String.format("%09d", i + 1);
            customers[i] = mobileNumber;
        }
        return customers;
    }

    public static void sendPromoDiscount(String[] input1, String[] input2) {
        for (int i = 0; i < input1.length; i++) {
            System.out.println("Sending Promo to " + input1[i]); // O(m)
        }
        System.out.println("It's Finished to send Promo to " + input1.length + " Customers"); // O(1)
        for (int i = 0; i < input2.length; i++) {
            System.out.println("Sending Discount to " + input2[i]); // O(n)
        }
        System.out.println("It's Finished to send Discount to " + input2.length + " Customers"); // O(1)
    }

    public static void main(String[] args) {
        String[] dataPromo = generateData(100000000);
        String[] dataDiscount = generateData(1000);
        sendPromoDiscount(dataPromo, dataDiscount);
    }
}

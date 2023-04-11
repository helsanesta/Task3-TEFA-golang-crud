class experiment11 {
    public static String[] generateData() {
        String number = "0821234567";
        String[] customers = new String[100];
        String mobileNumber;

        for (int i = 0; i < 100; i++) {
            if (i < 10) {
                mobileNumber = number + "0" + i;
            } else {
                mobileNumber = number + i;
            }
            customers[i] = mobileNumber;
        }

        return customers;
    }

    public static void sendPromoDiscount(String[] array) {
        for (int i = 0; i < array.length; i++) {
            System.out.println("Sending Promo to " + array[i]);
        }
        for (int i = 0; i < 10; i++) {
            System.out.println("Sending Discount to Chosen Customer " + (i + 1));
        }
    }

    public static void main(String[] args) {
        String[] customers = generateData();
        sendPromoDiscount(customers);
    }
}


class experiment10 {
    public static String[] telco = { "Telkom", "Indosat", "XL", "Verizon", "AT&T", "Nippon", "Vodafone", "Orange",
            "KDDI", "Telefonica", "T-Mobile" };

    public static void main(String[] args) {
        int company = (int) Math.round(Math.random() * 10);
        findCompany(telco, company);
    }

    public static void findCompany(String[] array, int input) {
        for (int i = 0; i < array.length; i++) {
            if (array[i] == array[input]) {
                System.out.println("Company Found " + array[input]);
                break;
            }
            System.out.println("Searching Company...");
        }
    }

}

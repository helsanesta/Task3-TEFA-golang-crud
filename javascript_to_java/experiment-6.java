import java.util.Arrays;

class experiment6 {
    public static void main(String[] args) {
        String address = "DKI Jakarta";
        String[] addresses = new String[100000];
        Arrays.fill(addresses, address);
        findAddress(addresses);
    }

    public static void findAddress(String[] addresses) {
        long tx = System.nanoTime();
        System.out.println("The Default Address is " + (addresses[0]));
        long ty = System.nanoTime();
        double performance = (ty - tx) / 1_000_000.0;
        System.out.println("The performance is " + performance + " ms");
    }
}
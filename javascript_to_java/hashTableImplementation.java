import java.util.ArrayList;

class HashTable {
    private ArrayList<String[]> data;

    public HashTable(int size) {
        this.data = new ArrayList<String[]>();
        for (int i = 0; i < size; i++) {
            this.data.add(null);
        }
    }

    private int hash(String key) {
        int hash = 0;
        for (int i = 0; i < key.length(); i++) {
            hash = (hash + key.charAt(i) * 1) % this.data.size();
        }
        return hash;
    }

    public void set(String key, String value) {
        int address = this.hash(key);
        if (this.data.get(address) == null) {
            this.data.set(address, new String[] { key, value });
        } else {
            this.data.get(address)[1] = value;
        }
    }

    public String get(String key) {
        int address = this.hash(key);
        String[] currentBucket = this.data.get(address);
        if (currentBucket != null) {
            for (int i = 0; i < currentBucket.length; i++) {
                if (currentBucket[i] != null && currentBucket[i].equals(key)) {
                    return currentBucket[i + 1];
                }
            }
        }
        return null;
    }

    public static void main(String[] args) {
        HashTable myHashTable = new HashTable(100);
        myHashTable.set("082124606606", "Rony Setyawan");
        System.out.println(myHashTable.get("082124606606"));
    }
}

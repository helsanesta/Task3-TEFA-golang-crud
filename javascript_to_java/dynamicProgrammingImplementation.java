import java.util.HashMap;
import java.util.Map;

class Fibonacci {
    private int counterPlainRecursive = 0;
    private int counterDynamicProgramming = 0;

    public int fibonacciPlainRecursive(int n) { // O(2^n)
        counterPlainRecursive++;
        if (n < 2) {
            return n;
        }
        return fibonacciPlainRecursive(n - 1) + fibonacciPlainRecursive(n - 2);
    }

    public int fibonacciDynamicProgramming(int n) { // O(n)
        Map<Integer, Integer> cache = new HashMap<>();
        return fib(n, cache);
    }

    private int fib(int n, Map<Integer, Integer> cache) {
        counterDynamicProgramming++;
        if (n < 2) {
            return n;
        } else if (cache.containsKey(n)) {
            return cache.get(n);
        } else {
            int result = fib(n - 1, cache) + fib(n - 2, cache);
            cache.put(n, result);
            return result;
        }
    }

    public static void main(String[] args) {
        Fibonacci fib = new Fibonacci();
        int n = 20;
        fib.fibonacciPlainRecursive(n);
        System.out.println("CounterPlainRecursive calculations for Plain Recursive: " + fib.counterPlainRecursive);
        fib.fibonacciDynamicProgramming(n);
        System.out.println(
                "CounterDynamicProgramming calculations for Dynamic Programming: " + fib.counterDynamicProgramming);
    }
}

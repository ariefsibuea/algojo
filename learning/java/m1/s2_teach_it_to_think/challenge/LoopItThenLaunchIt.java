public class LoopItThenLaunchIt {

    public static void main(String[] args) {
        for (int i = 1; i <= 10; i++) {
            System.out.println(i);
        }

        int x = 3;
        while (x >= 1) {
            System.out.println("Count down: " + x);
            x--;
        }
        System.out.println("Go!!!");
    }
}

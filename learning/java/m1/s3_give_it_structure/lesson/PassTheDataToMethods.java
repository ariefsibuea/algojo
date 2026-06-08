public class PassTheDataToMethods {

    public static int square(int x) {
        return x * x;
    }

    public static void main(String[] args) {
        int result = square(5);
        System.out.println("5 square is: " + result);
    }
}

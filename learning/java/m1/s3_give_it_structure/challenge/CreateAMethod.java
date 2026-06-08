public class CreateAMethod {

    public static void sayHiTo(String name) {
        System.out.println("Hi, " + name + "!");
    }

    public static int add(int a, int b) {
        return a + b;
    }

    public static void main(String[] args) {
        sayHiTo("Arief Sibuea");
        int result = add(3, 4);
        System.out.println("Sum is: " + result);
    }
}

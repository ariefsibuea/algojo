public class ReturnACustomGreeting {

    public static String greet(String name, int age) {
        return name + " is " + age + " years old.";
    }

    public static void main(String[] args) {
        String message = greet("Ali", 21);
        System.out.println(message);
    }
}

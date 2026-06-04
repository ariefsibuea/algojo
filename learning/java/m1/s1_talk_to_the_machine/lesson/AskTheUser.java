import java.util.Scanner;

public class AskTheUser {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.print("Enter your name: ");
        String name = input.nextLine();
        System.out.print("Enter your age: ");
        int age = Integer.parseInt(input.nextLine());
        System.out.print("Enter your address: ");
        String address = input.nextLine();

        System.out.println("Hi, " + name + "!");
        System.out.println("You are " + age + " years old.");
        System.out.println("You live in " + address + ".");

        input.close();
    }
}

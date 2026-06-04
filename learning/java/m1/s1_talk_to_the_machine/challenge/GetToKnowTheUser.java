import java.util.Scanner;

public class GetToKnowTheUser {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.print("Enter your name: ");
        String name = input.nextLine();

        System.out.print("Enter your favorite number: ");
        int number = input.nextInt();

        System.out.println("Hi, " + name + "!");
        System.out.println("Your favorite number is " + number + ".");

        input.close();
    }
}

import java.util.Scanner;

public class MakeDecisions {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.print("Enter your age: ");
        int age = input.nextInt();

        if (age >= 18) {
            System.out.println("You're an adult.");
        } else {
            System.out.println("You're a minor");
        }

        input.close();
    }
}

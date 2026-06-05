import java.util.Scanner;

public class BuildAQuiz {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("What's the capital of France?");
        System.out.println("1. Berlin");
        System.out.println("2. Madrid");
        System.out.println("3. Paris");
        System.out.print("Your choice: ");

        try {
            int choice = Integer.parseInt(input.nextLine());

            switch (choice) {
                case 1, 2:
                    System.out.println("Wrong answer.");
                    break;

                case 3:
                    System.out.println("Correct!!!");
                    break;

                default:
                    System.out.println("Select between 1 - 3.");
            }
        } catch (NumberFormatException e) {
            System.out.println("Invalid choice.");
        }

        input.close();
    }
}

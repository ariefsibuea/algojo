import java.util.Scanner;

public class PickAPath {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.print("Choose a number (1-3): ");
        int choice = input.nextInt();

        switch (choice) {
            case 1:
                System.out.println("Play game");
                break;

            case 2:
                System.out.println("Settings");
                break;

            case 3:
                System.out.println("Exit");
                break;

            default:
                System.out.println("Invalid choice");
        }

        input.close();
    }
}

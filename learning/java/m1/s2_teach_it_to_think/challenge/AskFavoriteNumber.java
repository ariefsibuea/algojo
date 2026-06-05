import java.util.Scanner;

public class AskFavoriteNumber {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.print("Enter your favorite number: ");
        int favoriteNumber = input.nextInt();

        if (favoriteNumber == 7) {
            System.out.println("Lucky!");
        } else if (favoriteNumber % 2 == 0) {
            System.out.println("Nice and even");
        } else {
            System.out.println("Interesting choice");
        }

        input.close();
    }
}

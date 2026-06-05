import java.util.Scanner;

public class TryCatch {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.print("Enter a number to divide 100 by: ");

        try {
            int number = Integer.parseInt(input.nextLine());
            int result = 100 / number;
            System.out.println("Result: " + result);
        } catch (NumberFormatException e) {
            System.out.println("Enter a valid number.");
        } catch (ArithmeticException e) {
            System.out.println("100 cannot be divided by the input number.");
        }

        input.close();
    }
}

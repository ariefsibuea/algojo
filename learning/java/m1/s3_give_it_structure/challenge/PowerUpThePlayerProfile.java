import java.util.Scanner;

public class PowerUpThePlayerProfile {

    public static class Player {
        String name;
        int level;
        double health;

        public Player(String name, int level, double health) {
            this.name = name;
            this.level = level;
            this.health = health;
        }

        public void displayProfile() {
            System.out.println("Name: " + this.name);
            System.out.println("Level: " + this.level);
            System.out.println("Health: " + this.health);
        }

        public void levelUp() {
            this.level += 1;
            this.health += 10;
        }
    }

    public static void main(String[] args) {
        try (Scanner input = new Scanner(System.in)) {
            System.out.print("Enter player name: ");
            String name = input.nextLine();

            System.out.print("Enter level: ");
            int level = Integer.parseInt(input.nextLine());

            System.out.print("Enter health: ");
            double health = Double.parseDouble(input.nextLine());
            ;

            Player hero = new Player(name, level, health);
            System.out.println("\n--- Player Created ---");
            hero.displayProfile();

            hero.levelUp();
            System.out.println("\n--- After Level Up ---");
            hero.displayProfile();
        } catch (NumberFormatException e) {
            System.out.println("Invalid input, try again ...");
        }
    }
}

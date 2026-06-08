public class AddAMethodToYourClass {

    public static class Player {
        String name;
        int score;

        public Player(String name, int score) {
            this.name = name;
            this.score = score;
        }

        public void display() {
            System.out.println(this.name + " scored " + this.score);
        }
    }

    public static void main(String[] args) {
        Player p1 = new Player("Aria", 50);
        Player p2 = new Player("Luca", 70);

        p1.display();
        p2.display();
    }
}

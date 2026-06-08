public class CreateObjects {

    public static class Player {
        String name;
        int score;

        public Player(String name, int score) {
            this.name = name;
            this.score = score;
        }
    }

    public static void main(String[] args) {
        Player p1 = new Player("Aria", 50);
        Player p2 = new Player("Luca", 70);

        System.out.println(p1.name + " scored " + p1.score);
        System.out.println(p2.name + " scored " + p2.score);
    }
}

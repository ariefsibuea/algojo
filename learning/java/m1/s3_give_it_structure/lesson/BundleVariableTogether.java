public class BundleVariableTogether {

    public static class Player {
        String name;
        int score;
    }

    public static void main(String[] args) {
        Player p = new Player();
        p.name = "Zara";
        p.score = 42;

        System.out.println(p.name + " scored " + p.score);
    }
}

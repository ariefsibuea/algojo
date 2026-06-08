public class BuildABookClass {

    public static class Book {
        String title;
        String author;
        int pages;
    }

    public static void main(String[] args) {
        Book book = new Book();
        book.title = "A Man Called John Doe";
        book.author = "Arief Sibuea";
        book.pages = 938;

        System.out.println(
                "A " + book.pages + "-page book with title \"" + book.title + "\" is written by " + book.author);
    }
}

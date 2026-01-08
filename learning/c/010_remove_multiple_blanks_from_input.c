#include <stdio.h>

int main()
{
    int c, nb;

    nb = 0;
    while ((c = getchar()) != EOF) {
        if (c == ' ' && nb > 0)
            ++nb;
        else if (c == ' ') {
            putchar(c);
            ++nb;
        } else {
            putchar(c);
            nb = 0;
        }
    }
}

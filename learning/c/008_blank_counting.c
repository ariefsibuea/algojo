#include <stdio.h>

int main()
{
    int c, nb;

    nb = 0;
    while ((c = getchar()) != EOF)
        if (c == ' ')
            ++nb;
    printf("number of blank: %d\n", nb);
}

#include <stdio.h>

int main()
{
    int c, nt;

    nt = 0;
    while ((c = getchar()) != EOF)
        if (c == '\t')
            ++nt;
    printf("number of tab: %d\n", nt);
}

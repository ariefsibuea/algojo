#include <stdio.h>

int main()
{
    // We must declare 'c' to be a type big enough to hold any value that 'getchar' returns. We can't use 'char' since
    // 'c' must be big enough to hold 'EOF' in addition to any possible 'char'. Therefore we use 'int'.
    int c;

    while ((c = getchar()) != EOF)
        putchar(c);

    printf("\nEOF reached, program is terminated. The value of EOF is: %d\n", c);
}

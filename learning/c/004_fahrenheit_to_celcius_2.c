#include <stdio.h>

// symbolic constants
#define LOWER 0
#define UPPER 300
#define STEP 20

int main()
{
    int fahr;

    // The 'for' loops is usually appropriate for loops in in which the initialization and increment are single
    // statements and logically related, since it is more compact than 'while' and it keeps the loop control statements
    // together in one place.
    for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP)
        printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
}

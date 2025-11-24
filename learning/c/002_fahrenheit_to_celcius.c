#include <stdio.h>

/* print Fahrenheit-Celcius table
 * for fahr = 0, 20, ..., 300
 */
int main()
{
    // we can direcly set the initial value like below
    // int fahr = 0, celcius = 0;
    // int lower = 0, upper = 300, step = 20;

    float fahr, celcius;
    float lower, upper, step;

    lower = 0;   /* lower limit of temperature scale */
    upper = 300; /* upper limit */
    step  = 20;  /* step size */

    fahr  = lower;
    while (fahr <= upper) {
        celcius = (5.0 / 9.0) * (fahr - 32.0);
        printf("%3.0f %6.1f\n", fahr, celcius);
        fahr = fahr + step;
    }

    return 0;
}

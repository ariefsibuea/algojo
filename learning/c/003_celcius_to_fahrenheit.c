#include <stdio.h>

int main()
{
    float fahr, celcius;
    float step, lower, upper;

    lower   = 0;
    upper   = 200;
    step    = 20;

    celcius = lower;
    while (celcius <= upper) {
        fahr = ((9.0 / 5.0) * celcius) + 32;
        printf("%3.0f %6.1f\n", celcius, fahr);
        celcius = celcius + step;
    }

    return 0;
}

#include<stdio.h>

int main() {
    int a = 21;
    int b = 10;
    int c;

    c = a+b;
    printf("Sum of 21 and 10 is %d\n", c);
    c = a-b;
    printf("Difference of 21 and 10 is %d\n", c);
    c = a*b;
    printf("Multiplication of 21 and 10 is %d\n", c);
    c = a/b;
    printf("Division of 21 and 10 is %d\n", c);
    c = a%b;
    printf("Remainder of 21 divided by 10 is %d\n", c);
    c = a++;
    printf("increment value of a by 1 using increment operator %d\n", a);
    c = a--;
    printf("decrement value of a by 1 using decrement operator %d\n", a);
}
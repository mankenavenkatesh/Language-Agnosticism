#include<iostream>
using namespace std;

int main() {

    int a = 21;
    int b = 10;
    int c;

    c = a+b;
    cout << "Sum of 21 and 10 is : " << c << endl;
    c = a-b;
    cout << "Difference of 21 and 10 is: " << c << endl;
    c = a * b;
    cout << "Multiplication of 21 and 10 is:" << c << endl;
    c = a / b;
    cout << "Division of 21 and 10 is:" << c << endl;
    c = a % b;
    cout << "Remainder of 21 divided by 10 is:" << c << endl;
    c = a++;
    cout << "Increment variable a by 1 using increment operator:" << a << endl;
    c = a--;
    cout << "Decrement variable a by 1 using decrement operator:" << a << endl;

    return 0;

}
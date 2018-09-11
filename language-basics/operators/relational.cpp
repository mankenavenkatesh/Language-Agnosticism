#include<iostream>
using namespace std;

int main() {
    int a = 21;
    int b = 10;
    int c;

    cout << "a=" << a << endl; 
    cout << "b=" << b << endl;
    if(a == b ) {
        cout << "a is equal to b" << endl;        
    }
    else {
        cout << "a is not equal to b" << endl;
    }

    if(a < b) {
        cout << "a is less than b" << endl;
    }
    else {
        cout << "a is not less than b" << endl;
    }

    if (a > b) {
        cout << "a is greater than b" << endl;
    }
    else {
        cout << "a is not greater than b" << endl;
    }

    a = 5;
    b = 20;
    cout << "a=" << a << endl;
    cout << "b=" << b << endl;
    
    if(a <= b){
        cout << "a is either less that or equal to b" << endl;
    }

    if(b >= a) {
        cout << "b is either greater than or equal to a" << endl;
    }
    return 0;

}
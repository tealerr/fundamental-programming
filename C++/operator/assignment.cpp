using namespace std;
#include<iostream>


void numberAssign();
void otherAssign();

int main(){
    numberAssign();
    otherAssign();

    return 0;
}

void numberAssign(){
    int x,y ;

    
    //Basic Assignment oparetor
    std::cout << "######### Basic Assignment oparetor #########" << std::endl;
    {
        x = 10;
        y = 5;
        x += y;
        std::cout << "x += y is " << x << std::endl;
    }
    

    {
        x = 10;
        y = 5;
        x -= y;
        std::cout << "x -= y is " << x << std::endl;
    }

    {
        x = 10;
        y = 5;
        x *= y;
        std::cout << "x *= y is " << x << std::endl;
    }
    
    {
        x = 10;
        y = 5;
        x /= y;
        std::cout << "x /= y is " << x << std::endl;
    }
    

    {
        x = 10;
        y = 5;
        x %= y;
        std::cout << "x %= y is " << x << std::endl;
    }

    
    //Shift Assignment Operators
    std::cout << "######### Shift Assignment Operators #########" << std::endl;
    {   
        x = 10;
        y = 5;
        x <<= 2; // Shift x left by 2 positions
        std::cout << "x after left shift: " << x << std::endl;
    }
    
    {   
        x = 10;
        y = 5;
        // Right shift assignment (x >>= y is equivalent to x = x >> y)
        x >>= 1; // Shift x right by 1 position
        std::cout << "x after right shift: " << x << std::endl;
    }
    
    //Bitwise Assignment Operators
    std::cout << "######### Bitwise Assignment Operators #########" << std::endl;
    {   
        x = 10;
        y = 5;
        x &= y; // Perform bitwise AND and assign the result to x
        std::cout << "x after bitwise AND: " << x << std::endl;
    }
    
    {   
        x = 10;
        y = 5;
        x |= y; // Perform bitwise OR and assign the result to x
        std::cout << "x after bitwise OR: " << x << std::endl;
    }

    {   
        x = 10;
        y = 5;
        x ^= y; // Perform bitwise XOR and assign the result to x
        std::cout << "x after bitwise XOR: " << x << std::endl;
    }

        std::cout << "\n" << std::endl;


}

void otherAssign(){
    int num = 10;
    double x = 3.14;
    float y = 5.0;
    string str1 = "Hello, ";
    char myChar = 'A';

    std::cout << "num: " << num << std::endl;
    std::cout << "x: " << x << std::endl;
    std::cout << "y: " << y << std::endl;
    std::cout << "str1: " << str1 << std::endl;
    std::cout << "myChar: " << myChar << std::endl;

}


using namespace std;
#include<iostream>

int numberLogic();
int decimalLogic();
void stringLogic();

int main(){
    numberLogic();
    decimalLogic();
    stringLogic();

    return 0;
}

int numberLogic(){
    int x = 5;
    int y = 3;

    if (x < 2 || y > 3) {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }

    if (x < 2 && y > 3) {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }

    if (!(x < 2)) {
        std::cout << "!(5 < 2) is True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }
        std::cout << "\n" << std::endl;

    return 0;
}

int decimalLogic(){
    double x = 5.25;
    double y = 3.5;

    if (x > 5.0 && y < 4.0) {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }

    if (x < 5.0 || y > 3.0) {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }

    if (!(x < 5.0 || y > 3.0)) {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }
    
    std::cout << "\n" << std::endl;

    return 0;
}

void stringLogic(){
    string myStr1 = "Hello";
    string myStr2 = "World";

    if (myStr1 == "hello" || myStr2 > "world") {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }

    if (myStr1 == "hello" && myStr2 > "world") {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }

    if (!(myStr1 == "hello")) {
        std::cout << "True" << std::endl;
    } else {
        std::cout << "False" << std::endl;
    }
    
    std::cout << "\n" << std::endl;

}


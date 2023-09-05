#include <iostream>

using namespace std;

void forLoop();
void doWhileLoop();
void whileLoop();
void breakAndCon();


int main() {
    forLoop();
    whileLoop();
    doWhileLoop();
    breakAndCon();
}

void forLoop() {
    for (int i = 0; i < 5; i++) {
        std::cout << "While count is " << i << std::endl;
    }
    std::cout << "\n";

}

void whileLoop() {
    int num = 1;
    while (num <= 5) {
        std::cout << "While count is " << num << std::endl;
        num++;
    }
    std::cout << "\n";


}

void doWhileLoop() {
    int num = 1;

    do {
        std::cout << "Do-while count is " << num << std::endl;
        num++;
    } while (num <= 5);
    std::cout << "\n";
}

void breakAndCon() {
    std::cout << "***** Use break \n";
    for (int i = 0; i < 5; i++) {
        if (i == 3) {
            std::cout << "pls break!!! \n";
            break;
        }
        std::cout << "Count is " << i << std::endl;
    }
    std::cout << "\n";

    std::cout << "***** Use continue \n";

    int num = 1;
    while (num <= 6) {
        if (num == 4) {
            std::cout << "skip this! \n";
            num++;
            continue;
        }

        std::cout << "Count is " << num << std::endl;
        num++;
    }

}
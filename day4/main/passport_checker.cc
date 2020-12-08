#include <fstream>
#include <iostream>
#include "passports.h"

int main(int argc, char* argv[]) {
    if (argc != 2) {
        std::cout << "Please specify path to inputs to check" << std::endl;
        return 1;
    }

    std::ifstream inputs(argv[1]);
    if (!inputs.is_open()) {
        std::cout << "error opening file" << std::endl;
        return 1;
    }

    day4::PassportReader t(&inputs);
    t.read();

    int validCount = 0;
    for (auto p : t.passports()) {
        if (p->valid()) {
            validCount++;
            continue;
        }
    }

    std::cout << validCount << std::endl;

    return 0;
}

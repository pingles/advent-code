#include <fstream>
#include <iostream>
#include "passwords.h"

using day2::InputTokeniser;
using day2::PasswordPtr;
using day2::PolicyPtr;

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

  InputTokeniser t(&inputs);
  PolicyPtr policy;
  PasswordPtr password;
  int validCount = 0;
  while (t.next(policy, password)) {
    if (password->satisfies(policy)) {
      validCount++;
      continue;
    }
  }

  std::cout << validCount << std::endl;

  return 0;
}

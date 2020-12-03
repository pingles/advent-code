#ifndef PASSWORDS_H
#define PASSWORDS_H

#include <iostream>
#include <memory>
#include <regex>
#include <sstream>
#include <string>
#include <utility>

namespace day2 {

class Policy {
 public:
  Policy(char ch, int min, int max) {
    character = ch;
    minMax = std::pair<int, int>(min, max);
  }
  char character;
  std::pair<int, int> minMax;
};

class Password {
 public:
  Password(std::string val) { value = val; }
  bool satisfies(const Policy* policy);

  std::string value;
};

typedef std::unique_ptr<Policy> PolicyPtr;
typedef std::unique_ptr<Password> PasswordPtr;

// reads a string line
class InputTokeniser {
 public:
  InputTokeniser(std::istringstream* input) {
    in = input;
    re = std::regex("(\\d+)\\-(\\d+)\\s([a-z]):\\s([a-z]+)");
  }

  bool next(PolicyPtr& policy, PasswordPtr& password);

 private:
  std::istringstream* in;
  std::regex re;
};

}  // namespace day2

#endif

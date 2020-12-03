#ifndef PASSWORDS_H
#define PASSWORDS_H

#include <iostream>
#include <istream>
#include <map>
#include <memory>
#include <regex>
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
typedef std::unique_ptr<Policy> PolicyPtr;

class Password {
 public:
  Password(std::string val) {
    value = val;
    index = std::make_unique<std::map<char, int>>();
    buildIndex();
  }
  bool satisfies(const PolicyPtr& policy);

  std::string value;

 private:
  std::unique_ptr<std::map<char, int>> index;
  int frequency(const char& c);
  void buildIndex();
};
typedef std::unique_ptr<Password> PasswordPtr;

// reads a string line
class InputTokeniser {
 public:
  InputTokeniser(std::istream* input) {
    in = input;
    re = std::regex("(\\d+)\\-(\\d+)\\s([a-z]):\\s([a-z]+)");
  }

  bool next(PolicyPtr& policy, PasswordPtr& password);

 private:
  std::istream* in;
  std::regex re;
};

}  // namespace day2

#endif

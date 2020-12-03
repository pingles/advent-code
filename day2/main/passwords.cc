#include "passwords.h"
#include <memory>
#include <regex>

namespace day2 {

bool InputTokeniser::next(PolicyPtr& policy, PasswordPtr& password) {
  // read line
  std::string s;
  std::getline(*in, s);

  std::smatch matches;
  std::regex_match(s, matches, re);

  if (matches.size() != 5) {
    return false;
  }

  int min = std::stoi(matches[1].str());
  int max = std::stoi(matches[2].str());
  char character = matches[3].str().at(0);

  policy = std::make_unique<Policy>(character, min, max);
  password = std::make_unique<Password>(matches[4].str());

  return true;
}

}  // namespace day2

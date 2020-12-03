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

void Password::buildIndex() {
  for (const char& c : value) {
    if (index->find(c) == index->end()) {
      (*index)[c] = 1;
    } else {
      (*index)[c]++;
    }
  }
}

int Password::frequency(const char& c) {
  auto val = index->find(c);
  if (val == index->end()) {
    return 0;
  }
  return val->second;
}

// determines if contents of password matches policy
// expectation
bool Password::satisfies(const PolicyPtr& policy) {
  auto actual = frequency(policy->character);
  if (actual == 0) {
    return false;
  }

  if (actual < policy->minMax.first || actual > policy->minMax.second) {
    return false;
  }

  return true;
}

}  // namespace day2

#include "passports.h"
#include <forward_list>
#include <iostream>
#include <memory>
#include <regex>
#include <sstream>
#include <string>
#include <vector>

namespace day4 {

bool match_re(const std::string v, std::string pattern) {
  std::regex re(pattern);
  return std::regex_match(v, re);
}

// checks if v has x digits
bool digits(const std::string v, int digits) {
  std::ostringstream expr;
  expr << "[\\d]{" << digits << "}";
  return match_re(v, expr.str());
}

bool range(const std::string v, int min, int max) {
  auto n = std::stoi(v);
  return n >= min && n <= max;
}

bool Passport::valid_data(const std::string k, const std::string v) {
  if (k == "byr") {
    return digits(v, 4) && range(v, 1920, 2002);
  }
  if (k == "iyr") {
    return digits(v, 4) && range(v, 2010, 2020);
  }
  if (k == "eyr") {
    return digits(v, 4) && range(v, 2020, 2030);
  }
  if (k == "hcl") {
    return match_re(v, "^\\#[0-9a-f]{6}$");
  }
  if (k == "ecl") {
    return match_re(v, "^(amb|blu|brn|gry|grn|hzl|oth)$");
  }
  if (k == "pid") {
    return match_re(v, "^[0-9]{9}$");
  }
  if (k == "hgt") {
    if (match_re(v, "\\d+cm$")) {
      return range(v, 150, 193);
    } else if (match_re(v, "\\d+in$")) {
      return range(v, 59, 76);
    }
  }

  return false;
}

void Passport::add_data(const std::string k, const std::string v) {
  if (valid_data(k, v)) {
    required->erase(k);
  }
}

bool Passport::valid() {
  return required->empty();
}

void Passport::message() {
  for (auto x = required->begin(); x != required->end(); x++) {
    std::cout << *x << std::endl;
  }
}

std::shared_ptr<Passport> passport_from_data(std::forward_list<std::string> data) {
  auto passport = std::make_shared<Passport>();
  for (auto l : data) {
    static std::regex re("((byr|iyr|eyr|hgt|hcl|ecl|pid|cid)):([\\S]+)");
    std::sregex_iterator next(l.begin(), l.end(), re);
    std::sregex_iterator end;
    while (next != end) {
      std::smatch match = *next;
      auto key = match[2];
      auto value = match[3];
      passport->add_data(key, value);
      next++;
    }
  }
  return passport;
}

void PassportReader::read() {
  std::forward_list<std::string> data;

  std::string line;
  while (std::getline(*inputStream, line)) {
    if (line.length() > 0 && line[0] != '\n') {
      // push data
      data.push_front(line);
    } else {
      auto passport = passport_from_data(data);
      _passports.push_back(passport);
      data.clear();  // ready for the next record
    }
  }

  // did we have an entry on the final line of the input
  if (!data.empty()) {
    auto passport = passport_from_data(data);
    _passports.push_back(passport);
  }
}

std::vector<std::shared_ptr<Passport>>& PassportReader::passports() {
  return _passports;
}

}  // namespace day4

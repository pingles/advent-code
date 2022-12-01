#include <istream>
#include <memory>
#include <set>
#include <string>
#include <vector>
#include <regex>

namespace day4 {

class Passport {
 public:
  Passport() {
    required = std::make_unique<std::set<std::string>>();
    required->insert("byr");
    required->insert("iyr");
    required->insert("eyr");
    required->insert("hgt");
    required->insert("hcl");
    required->insert("ecl");
    required->insert("pid");
  }
  void add_data(const std::string k, const std::string v);
  void message();
  bool valid();
  static bool valid_data(const std::string k, const std::string v);

 private:
  std::unique_ptr<std::set<std::string>> required;
};

// reads from an input stream, uses a state
// machine to move between states
class PassportReader {
 public:
  PassportReader(std::istream* input) {
    inputStream = input;
    passports() = std::vector<std::shared_ptr<Passport>>{};
  };
  void read();  // processes the input
  std::vector<std::shared_ptr<Passport>>& passports();

 private:
  std::istream* inputStream;
  std::vector<std::shared_ptr<Passport>> _passports;
};

}  // namespace day4

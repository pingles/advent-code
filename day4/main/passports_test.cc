#include "passports.h"
#include <sstream>
#include "gtest/gtest.h"

namespace day4 {
TEST(passport_data_validation, validate_field_data) {
  EXPECT_EQ(false, Passport::valid_data("byr", "001920"));
  EXPECT_EQ(false, Passport::valid_data("byr", "1900"));
  EXPECT_EQ(true, Passport::valid_data("byr", "1920"));

  EXPECT_EQ(false, Passport::valid_data("iyr", "002019"));
  EXPECT_EQ(false, Passport::valid_data("iyr", "2009"));
  EXPECT_EQ(true, Passport::valid_data("iyr", "2010"));

  EXPECT_EQ(false, Passport::valid_data("eyr", "002030"));
  EXPECT_EQ(false, Passport::valid_data("eyr", "2010"));
  EXPECT_EQ(true, Passport::valid_data("eyr", "2030"));

  EXPECT_EQ(false, Passport::valid_data("hcl", "#9999999"));
  EXPECT_EQ(false, Passport::valid_data("hcl", "#aaaaaa "));
  EXPECT_EQ(true, Passport::valid_data("hcl", "#aaaaaa"));

  EXPECT_EQ(false, Passport::valid_data("ecl", "foo"));
  EXPECT_EQ(true, Passport::valid_data("ecl", "amb"));

  EXPECT_EQ(true, Passport::valid_data("pid", "123456789"));
  EXPECT_EQ(true, Passport::valid_data("pid", "000000000"));
  EXPECT_EQ(false, Passport::valid_data("pid", "1234567890"));
  EXPECT_EQ(false, Passport::valid_data("pid", " 123456789"));

  EXPECT_EQ(true, Passport::valid_data("hgt", "150cm"));
  EXPECT_EQ(false, Passport::valid_data("hgt", "149cm"));
  EXPECT_EQ(false, Passport::valid_data("hgt", "78in"));
  EXPECT_EQ(true, Passport::valid_data("hgt", "76in"));
}

TEST(passport_data_validation, invalid_records) {
  std::istringstream sample(
      "eyr:1972 cid:100\n"
      "hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n"
      "\n"
      "iyr:2019\n"
      "hcl:#602927 eyr:1967 hgt:170cm\n"
      "ecl:grn pid:012533040 byr:1946\n"
      "\n"
      "hcl:dab227 iyr:2012\n"
      "ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n"
      "\n"
      "hgt:59cm ecl:zzz\n"
      "eyr:2038 hcl:74454a iyr:2023\n"
      "pid:3556412378 byr:2007\n\n");
  PassportReader r(&sample);
  r.read();

  for (int i = 0; i < 4; i++) {
    EXPECT_EQ(false, r.passports().at(i)->valid());
  }
}

TEST(passport_data_validation, valid_records) {
  std::istringstream sample(
      "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\n"
      "hcl:#623a2f\n"
      "\n"
      "eyr:2029 ecl:blu cid:129 byr:1989\n"
      "iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n"
      "\n"
      "hcl:#888785\n"
      "hgt:164cm byr:2001 iyr:2015 cid:88\n"
      "pid:545766238 ecl:hzl\n"
      "eyr:2022\n"
      "\n"
      "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719\n");
  PassportReader r(&sample);
  r.read();

  for (int i = 0; i < 4; i++) {
    EXPECT_EQ(true, r.passports().at(i)->valid());
  }
}

TEST(passports, valid_record) {
  std::istringstream sample("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\n");
  PassportReader r(&sample);
  r.read();

  EXPECT_EQ(1, r.passports().size());
  EXPECT_EQ(true, r.passports().at(0)->valid());
}

TEST(passports, full_sample) {
  std::istringstream sample(
      "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n"
      "byr:1937 iyr:2017 cid:147 hgt:183cm\n"
      "\n"
      "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n"
      "hcl:#cfa07d byr:1929\n"
      "\n"
      "hcl:#ae17e1 iyr:2013\n"
      "eyr:2024\n"
      "ecl:brn pid:760753108 byr:1931\n"
      "hgt:179cm\n"
      "\n"
      "hcl:#cfa07d eyr:2025 pid:166559648\n"
      "iyr:2011 ecl:brn hgt:59in\n\n");
  PassportReader r(&sample);
  r.read();

  EXPECT_EQ(4, r.passports().size());
  EXPECT_EQ(true, r.passports().at(0)->valid());
  EXPECT_EQ(false, r.passports().at(1)->valid());
  EXPECT_EQ(true, r.passports().at(2)->valid());
  EXPECT_EQ(false, r.passports().at(3)->valid());
}

}  // namespace day4

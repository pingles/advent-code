#include "passwords.h"
#include <sstream>
#include "gtest/gtest.h"

namespace day2 {

TEST(passwords_test, empty) {
  std::istringstream sample("");
  InputTokeniser t(&sample);

  PolicyPtr po;
  PasswordPtr pa;
  ASSERT_EQ(false, t.next(po, pa));
}

TEST(passwords_test, simple) {
  std::istringstream sample("1-3 a: abb");
  InputTokeniser t(&sample);

  PolicyPtr po;
  PasswordPtr pa;
  ASSERT_EQ(true, t.next(po, pa));
  ASSERT_EQ(1, po->minMax.first);
  ASSERT_EQ(3, po->minMax.second);
  ASSERT_EQ("abb", pa->value);
}

}  // namespace day2

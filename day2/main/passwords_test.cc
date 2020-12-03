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

TEST(passwords_test, part2_data) {
  std::istringstream sample("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc");
  InputTokeniser t(&sample);

  PolicyPtr po;
  PasswordPtr pa;
  t.next(po, pa);

  ASSERT_EQ(true, pa->satisfies(po)); // position 1 contains a and position 3 does not.

  t.next(po, pa);
  ASSERT_EQ(false, pa->satisfies(po)); // neither position 1 nor position 3 contains b.

  t.next(po, pa);
  ASSERT_EQ(false, pa->satisfies(po)); // both position 2 and position 9 contain c.
}

}  // namespace day2

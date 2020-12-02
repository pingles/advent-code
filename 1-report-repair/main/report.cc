#include "report.h"
#include <set>

// Our input will be a file containing a list of unsorted numbers.
// We need to find two entries that sum to 2020, multiply them, and return the
// value.
//
// Simple:
// We compare each number against each other number, this would be O(n^2)
//
// Complex:
// We know our target is 2020, so we could put all numbers into a set.
// Inserts and find should be O(1), so overall solution is O(n) as we'll
// need to iterate over input to insert and find.
//
// PART 2
// We need to find 3 numbers.
//
// for i = 0; i < n
//   for j = i + 1; j < n
//      if exists(2020 - inputs[i] - inputs[j])
//         found
//
// this is O(n^2). is there a faster way?
// if we have all the numbers ordered descending
// we know that 2020 > inputs[i] > inputs[j]
// so if we have the numbers ordered, can only iterate
// the numbers smaller than inputs[i]

using std::tuple;

tuple<int, int> target_pairs(const std::unordered_set<int> inputs,
                                  const int target) {
  // find pairs
  for (const int a : inputs) {
    const int b = target - a;
    if (inputs.find(b) != inputs.end()) {
      return tuple<int, int>(a, b);
    }
  }
  return tuple<int, int>();
}

tuple<int, int, int> target_trios(
    const std::set<int> inputs, const int target) {
  auto a = inputs.begin();
  while (a != inputs.end()) {
    // inner loop starts at outer+1
    auto b = a++;
    while (b != inputs.end()) {
      // 2020 > a > b
      // check whether val = 2020 - a - b is present
      auto it = target - *a - *b;
      if (inputs.find(it) != inputs.end()) {
        return tuple<int, int, int>(*a, *b, it);
      }
      b++;
    }

    a++;
  }

  return tuple<int, int, int>();
}

#ifndef REPORT_H
#define REPORT_H

#import <set>
#import <tuple>
#import <unordered_set>

namespace day1 {
// finds pairs of numbers that sum to target
std::tuple<int, int> target_pairs(const std::unordered_set<int> values,
                                  const int target);
// finds trios of numbers that sum to target
std::tuple<int, int, int> target_trios(const std::set<int> values,
                                       const int target);
}  // namespace day1

#endif

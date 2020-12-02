#ifndef REPORT_H
#define REPORT_H

#import <unordered_set>
#import <set>
#import <tuple>

// finds pairs of numbers that sum to target
std::tuple<int,int> target_pairs(const std::unordered_set<int> values, const int target);
// finds trios of numbers that sum to target
std::tuple<int,int,int> target_trios(const std::set<int,std::greater<int>> values, const int target);

#endif

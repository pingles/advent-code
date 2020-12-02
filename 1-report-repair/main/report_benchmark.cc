#include <benchmark/benchmark.h>
#include <iostream>
#include <random>
#include "report.h"

std::set<int> randomSet(const int size) {
  std::set<int> x;

  std::default_random_engine generator;
  std::uniform_int_distribution<int> distribution(1, 32768);

  for (int i = 0; i < size; i++) {
    x.insert(distribution(generator));
  }

  return x;
}

static void BM_target_trios(benchmark::State& state) {
  while (state.KeepRunning()) {
    state.PauseTiming();
    auto sample = randomSet(state.range(0));
    state.ResumeTiming();
    target_trios(sample, 32768);
    state.SetComplexityN(state.range(0));
  }
}

BENCHMARK(BM_target_trios)->RangeMultiplier(2)->Range(8, 8 << 12)->Complexity();
BENCHMARK_MAIN();

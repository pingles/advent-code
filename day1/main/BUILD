cc_library(
    name = "report",
    srcs = ["report.cc"],
    hdrs = ["report.h"],
    visibility = ["//visibility:public"],
)

cc_binary(
    name = "report_bench",
    srcs = ["report_benchmark.cc"],
    deps = ["//day1/main:report",
        "@benchmark//:benchmark_main"],
)

cc_test(
    name = "report_test",
    srcs = ["report_test.cc"],
    deps = ["//day1/main:report",
           "@googletest//:gtest_main"],
)
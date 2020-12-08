cc_library(
    name = "passports",
    srcs = ["passports.cc"],
    hdrs = ["passports.h"],
    visibility = ["//visibility:public"],
)

cc_test(
    name = "passports_test",
    srcs = ["passports_test.cc"],
    deps = ["//day4/main:passports",
           "@googletest//:gtest_main"],
)

cc_binary(
    name = "passport_checker",
    srcs = ["passport_checker.cc"],
    deps = ["//day4/main:passports"],
)
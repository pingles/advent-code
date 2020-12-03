cc_library(
    name = "passwords",
    srcs = ["passwords.cc"],
    hdrs = ["passwords.h"],
    visibility = ["//visibility:public"],
)

cc_test(
    name = "passwords_test",
    srcs = ["passwords_test.cc"],
    deps = ["//day2/main:passwords",
           "@googletest//:gtest_main"],
)
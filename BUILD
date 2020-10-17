load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/jmeekhof/epiweek
gazelle(name = "gazelle")

go_library(
    name = "epiweek",
    srcs = ["epiweek.go"],
    importpath = "github.com/jmeekhof/epiweek",
    visibility = ["//visibility:public"],
)

go_test(
    name = "epiweek_test",
    size = "small",
    srcs = ["epiweek_test.go"],
    embed = [":epiweek"],
)

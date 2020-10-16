load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/jmeekhof/epiweek
gazelle(name = "gazelle")

go_library(
    name = "epiweek",
    srcs = ["epiweek.go"],
    importpath = "github.com/jmeekhof/epiweek",
    visibility = ["//visibility:public"],
)

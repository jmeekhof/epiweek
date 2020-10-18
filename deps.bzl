load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_mattn_goveralls",
        strip_prefix = "goveralls-0.0.7",
        importpath = "github.com/mattn/goveralls",
        urls = ["https://github.com/mattn/goveralls/archive/v0.0.7.zip"],
        type = "zip",
    )

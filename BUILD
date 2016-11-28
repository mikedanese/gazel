package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
    "go_test",
    "cgo_library",
)

go_binary(
    name = "gazel",
    srcs = [
        "config.go",
        "gazel.go",
    ],
    tags = ["automanaged"],
    deps = [
        "///third_party/go/path/filepath:go_default_library",
        "//vendor:github.com/bazelbuild/buildifier/core",
        "//vendor:github.com/golang/glog",
    ],
)

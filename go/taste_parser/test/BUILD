load("@rules_go//go:def.bzl", "go_test")

go_test(
    name = "taste_parser",
    srcs = ["parser_test.go"],
    deps = [
        "//go/shared",
        "//go/taste_parser",
        "//go/taste_tokenizer"
    ],
    visibility = ["//visibility:private"],
)

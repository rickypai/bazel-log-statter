load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.0/rules_go-0.19.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.0/rules_go-0.19.0.tar.gz",
    ],
    sha256 = "9fb16af4d4836c8222142e54c9efa0bb5fc562ffc893ce2abeac3e25daead144",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
    ],
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
)

BAZEL_BUILD_TOOLS_SHA = "4bcdbd1064fcc48180fa30400e39f7a940fdb8f9"

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-"+BAZEL_BUILD_TOOLS_SHA,
    url = "https://github.com/bazelbuild/buildtools/archive/"+BAZEL_BUILD_TOOLS_SHA+".zip",
)

ATLASSIAN_BAZEL_TOOLS_SHA = "02472d814d1867de2ad75801760732c9e595c9d9"

http_archive(
    name = "com_github_atlassian_bazel_tools",
    strip_prefix = "bazel-tools-"+ATLASSIAN_BAZEL_TOOLS_SHA,
    urls = ["https://github.com/atlassian/bazel-tools/archive/"+ATLASSIAN_BAZEL_TOOLS_SHA+".zip"],
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "f40c92d1b30ff758a66aba7578039cbf959aea62",
    remote = "https://github.com/bazelbuild/rules_docker.git",
)

git_repository(
    name = "io_bazel_rules_k8s",
    commit = "87de67e50c1135a6b6bd0f1401f004bf7c711c42",
    remote = "https://github.com/bazelbuild/rules_k8s.git",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

load("@com_github_atlassian_bazel_tools//goimports:deps.bzl", "goimports_dependencies")

goimports_dependencies()

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

k8s_repositories()

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    commit = "34c6fa2dc70986bccbbffcc6130f6920a924b075", # Mar 4, 2019
)

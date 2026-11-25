# pulseten-go-libs

This repository contains prebuilt static libraries of
[PulseTen](https://github.com/leviathan0992/pulseten), an embedded time-series
database built on DuckDB. The PulseTen Go package links them, so that it builds
without compiling PulseTen or DuckDB.

## Installation

Install the PulseTen Go package. It downloads the matching libraries from here:

```sh
go get github.com/leviathan0992/pulseten/go
```

`go build` needs a C/C++ compiler to link the libraries, such as `g++` on Linux
or the Xcode command line tools on macOS. PulseTen and DuckDB are not compiled
again.

## Platforms

Each directory is a separate Go module. A build downloads only the module of its
target platform.

| Module | `GOOS/GOARCH` | Tested on |
| --- | --- | --- |
| `darwin-arm64` | `darwin/arm64` | macOS 15; requires macOS 12 or newer |
| `linux-amd64` | `linux/amd64` | Ubuntu 22.04 and 24.04, Debian 12 |
| `linux-arm64` | `linux/arm64` | Ubuntu 22.04 and 24.04, Debian 12 |
| `linux-amd64-musl` | `linux/amd64` | Alpine 3.20 and 3.22 |
| `linux-arm64-musl` | `linux/arm64` | Alpine 3.20 and 3.22 |

`GOOS` and `GOARCH` are the same for glibc and musl, so the musl modules are
selected with a build tag:

```sh
go build -tags pulseten_use_musl ./...
```

Other platforms require a
[build from source](https://github.com/leviathan0992/pulseten#building-from-source).

## Contents

| File | Description |
| --- | --- |
| `libpulseten.a` | PulseTen |
| `libduckdb.a` | DuckDB; on Linux split into numbered parts, `libduckdb_1.a` and so on, to stay below GitHub's file size limit |
| `libduckdb_core_functions.a` | DuckDB's built-in functions |
| `lib.go` | cgo flags |

## Releases

Each PulseTen release publishes all modules with its own version number. Go
requires the directory of a module in its tag, so the release has one tag per
module:

```
v0.1.0  ->  darwin-arm64/v0.1.0
            linux-amd64/v0.1.0
            linux-arm64/v0.1.0
            linux-amd64-musl/v0.1.0
            linux-arm64-musl/v0.1.0
```

The [release workflow](https://github.com/leviathan0992/pulseten/blob/main/.github/workflows/release.yml)
of PulseTen builds and tests the libraries and creates the tags. A published tag
is never changed, because the Go module proxy keeps every version permanently.

## License

PulseTen is licensed under the [Apache License 2.0](LICENSE). The licenses of
DuckDB and the other bundled components are listed in
[THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).

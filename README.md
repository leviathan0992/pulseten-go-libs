# pulseten-go-libs

Prebuilt static libraries for the [PulseTen](https://github.com/leviathan0992/pulseten)
Go package, one Go module per platform.

Do not depend on this repository directly. Import the PulseTen package, which
selects the module for your platform:

```sh
go get github.com/leviathan0992/pulseten/go
```

## Contents

Each platform directory is a Go module holding the libraries for one platform.
The musl modules are selected with the `pulseten_use_musl` build tag:

| Directory | Platform |
| --- | --- |
| `darwin-arm64` | macOS on Apple silicon |
| `linux-amd64` | Linux on x86-64, glibc |
| `linux-amd64-musl` | Linux on x86-64, musl |
| `linux-arm64` | Linux on ARM64, glibc |
| `linux-arm64-musl` | Linux on ARM64, musl |

A module contains `libpulseten.a` with PulseTen itself, `libduckdb.a` and
`libduckdb_core_functions.a` with the DuckDB build it uses, and a small cgo file
that links them. On Linux `libduckdb.a` is split into parts to stay under the
file size limit of GitHub. Building still requires cgo and a C toolchain with
the C++ standard library, but not CMake, a C++ build of DuckDB or pkg-config.

## Releases

The PulseTen release workflow builds, tests and commits the libraries here,
then tags every platform module with the PulseTen version, for example
`darwin-arm64/v0.1.0`. Do not edit or retag this repository by hand: the Go
module proxy caches every published version permanently.

## Licenses

PulseTen is licensed under the Apache License 2.0; see [LICENSE](LICENSE). The
libraries also contain DuckDB and the third-party components it links, under
the licenses in [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).

// Copyright (c) 2026, The PulseTen Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux && arm64

// Package linux_arm64_musl links the prebuilt PulseTen and DuckDB libraries for
// Linux on ARM64 with musl. Import github.com/leviathan0992/pulseten/go instead.
package linux_arm64_musl

// #cgo LDFLAGS: -L${SRCDIR} -Wl,--start-group -lpulseten -lduckdb_core_functions -lduckdb_1 -lduckdb_2 -Wl,--end-group -lstdc++ -lm -ldl -lpthread
import "C"

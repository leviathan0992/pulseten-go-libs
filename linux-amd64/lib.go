// Copyright (c) 2026, The PulseTen Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux && amd64

// Package linux_amd64 links the prebuilt PulseTen and DuckDB libraries for
// Linux on x86-64 with glibc. Import github.com/leviathan0992/pulseten/go instead.
package linux_amd64

// #cgo LDFLAGS: -L${SRCDIR} -Wl,--start-group -lpulseten -lduckdb_core_functions -lduckdb_1 -lduckdb_2 -Wl,--end-group -lstdc++ -lm -ldl -lpthread
import "C"

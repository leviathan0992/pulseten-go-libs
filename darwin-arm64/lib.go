// Copyright (c) 2026, The PulseTen Authors
// SPDX-License-Identifier: Apache-2.0

//go:build darwin && arm64

// Package darwin_arm64 links the prebuilt PulseTen and DuckDB libraries for
// macOS on Apple silicon. Import github.com/leviathan0992/pulseten/go instead.
package darwin_arm64

// #cgo LDFLAGS: -L${SRCDIR} -lpulseten -lduckdb_core_functions -lduckdb -lc++
import "C"

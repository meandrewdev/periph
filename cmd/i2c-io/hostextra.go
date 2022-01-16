// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

//go:build periphextra
// +build periphextra

package main

import (
	"github.com/meandrewdev/periph"
	"periph.io/x/extra/hostextra"
)

func hostInit() (*periph.State, error) {
	return hostextra.Init()
}

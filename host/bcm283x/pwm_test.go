// Copyright 2017 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package bcm283x

import (
	"testing"

	"github.com/meandrewdev/periph/conn/physic"
)

func TestPWMMap(t *testing.T) {
	defer reset()
	p := pwmMap{}
	p.reset()
	if _, err := setPWMClockSource(); err == nil {
		t.Fatal("pwmMemory is nil")
	}
	drvDMA.pwmMemory = &p
	if _, err := setPWMClockSource(); err == nil {
		t.Fatal("clockMemory is nil")
	}
	drvDMA.clockMemory = &clockMap{}
	drvDMA.pwmBaseFreq = 25 * physic.MegaHertz
	drvDMA.pwmDMAFreq = 200 * physic.KiloHertz
	if _, err := setPWMClockSource(); err == nil {
		t.Fatal("can't write to clock register")
	}
}

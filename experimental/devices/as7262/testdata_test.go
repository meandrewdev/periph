// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package as7262

import (
	"time"

	"github.com/meandrewdev/periph/conn/i2c/i2ctest"
	"github.com/meandrewdev/periph/conn/physic"
)

// Expected response from sensorTestCaseValidRead or
// sensorTestCaseInteruptValidRead.
var validSpectrum = Spectrum{
	Bands: []Band{
		{450 * physic.NanoMetre, 0.15625, 43707, "V"},
		{500 * physic.NanoMetre, 0.15625, 43707, "B"},
		{550 * physic.NanoMetre, 0.15625, 43707, "G"},
		{570 * physic.NanoMetre, 0.15625, 43707, "Y"},
		{600 * physic.NanoMetre, 0.15625, 43707, "O"},
		{650 * physic.NanoMetre, 0.15625, 43707, "R"},
	},
	SensorTemperature: physic.ZeroCelsius,
	LedDrive:          physic.MilliAmpere * 100,
	Integration:       2800 * time.Microsecond,
}

// Sequence of i2c traffic that yeilds validSpectrum.
var sensorTestCaseValidRead = []i2ctest.IO{
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x85}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x01}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x87}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x38}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x84}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0c}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x04}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x02}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x87}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x00}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x08}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x09}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0a}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0b}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0c}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0d}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0e}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0f}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x10}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x11}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x12}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x13}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x14}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x15}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x16}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x17}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x18}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x19}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1a}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1b}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1c}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1d}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1e}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1f}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x20}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x21}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x22}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x23}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x24}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x25}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x26}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x27}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x28}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x29}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x2a}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x2b}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x06}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
}

// Same as sensorTestCaseValidRead but omitting polling for data ready.
var sensorTestCaseInteruptValidRead = []i2ctest.IO{
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x85}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x01}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x87}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x38}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x84}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0c}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x87}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x00}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x08}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x09}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0a}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0b}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0c}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0d}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0e}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x0f}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x10}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x11}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x12}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xAA}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x13}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0xBB}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x14}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x15}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x16}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x17}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x18}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x19}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1a}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1b}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1c}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1d}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1e}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x1f}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x20}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x21}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x22}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x23}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x24}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x25}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x26}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x27}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x28}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x3e}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x29}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x20}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x2a}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x2b}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{writeReg, 0x06}, R: []byte{}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x01}},
	{Addr: 0x49, W: []byte{readReg}, R: []byte{0x00}},
	{Addr: 0x49, W: []byte{statusReg}, R: []byte{0x00}},
}

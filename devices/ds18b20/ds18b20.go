// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package ds18b20

import (
	"errors"
	"time"

	"github.com/meandrewdev/periph/conn"
	"github.com/meandrewdev/periph/conn/onewire"
	"github.com/meandrewdev/periph/conn/physic"
)

// ConvertAll performs a conversion on all DS18B20 devices on the bus.
//
// During the conversion it places the bus in strong pull-up mode to power
// parasitic devices and returns when the conversions have completed. This time
// period is determined by the maximum resolution of all devices on the bus and
// must be provided.
//
// ConvertAll uses time.Sleep to wait for the conversion to finish, which takes
// from 94ms to 752ms.
func ConvertAll(o onewire.Bus, maxResolutionBits int) error {
	if maxResolutionBits < 9 || maxResolutionBits > 12 {
		return errors.New("ds18b20: invalid maxResolutionBits")
	}
	if err := o.Tx([]byte{0xcc, 0x44}, nil, onewire.StrongPullup); err != nil {
		return err
	}
	conversionSleep(maxResolutionBits)
	return nil
}

// New returns an object that communicates over 1-wire to the DS18B20 sensor
// with the specified 64-bit address.
//
// resolutionBits must be in the range 9..12 and determines how many bits of
// precision the readings have. The resolution affects the conversion time:
// 9bits:94ms, 10bits:188ms, 11bits:375ms, 12bits:750ms.
//
// A resolution of 10 bits corresponds to 0.25C and tends to be a good
// compromise between conversion time and the device's inherent accuracy of
// +/-0.5C.
func New(o onewire.Bus, addr onewire.Address, resolutionBits int) (*Dev, error) {
	if resolutionBits < 9 || resolutionBits > 12 {
		return nil, errors.New("ds18b20: invalid resolutionBits")
	}

	d := &Dev{onewire: onewire.Dev{Bus: o, Addr: addr}, resolution: resolutionBits}

	// Start by reading the scratchpad memory, this will tell us whether we can
	// talk to the device correctly and also how it's configured.
	spad, err := d.readScratchpad()
	if err != nil {
		return nil, err
	}

	// Change the resolution, if necessary (datasheet p.6).
	if int(spad[4]>>5) != resolutionBits-9 {
		// Set the value in the configuration register.
		if err := d.onewire.Tx([]byte{0x4e, 0, 0, byte((resolutionBits-9)<<5) | 0x1f}, nil); err != nil {
			return nil, err
		}
		// Copy the scratchpad to EEPROM to save the values.
		if err := d.onewire.TxPower([]byte{0x48}, nil); err != nil {
			return nil, err
		}
		// Wait for the write to complete.
		sleep(10 * time.Millisecond)
	}

	return d, nil
}

// Dev is a handle to a Dallas Semi / Maxim DS18B20 temperature sensor on a
// 1-wire bus.
type Dev struct {
	onewire    onewire.Dev // device on 1-wire bus
	resolution int         // resolution in bits (9..12)
}

func (d *Dev) String() string {
	return "DS18B20{" + d.onewire.String() + "}"
}

// Halt implements conn.Resource.
func (d *Dev) Halt() error {
	return nil
}

// Sense implements physic.SenseEnv.
func (d *Dev) Sense(e *physic.Env) error {
	if err := d.onewire.TxPower([]byte{0x44}, nil); err != nil {
		return err
	}
	conversionSleep(d.resolution)
	t, err := d.LastTemp()
	if err != nil {
		return err
	}
	e.Temperature = t
	return nil
}

// SenseContinuous implements physic.SenseEnv.
func (d *Dev) SenseContinuous(time.Duration) (<-chan physic.Env, error) {
	// TODO(maruel): Manually poll in a loop via time.NewTicker.
	return nil, errors.New("ds18b20: not implemented")
}

// Precision implements physic.SenseEnv.
func (d *Dev) Precision(e *physic.Env) {
	e.Temperature = physic.Kelvin / 16
}

// LastTemp reads the temperature resulting from the last conversion from the
// device.
//
// It is useful in combination with ConvertAll.
func (d *Dev) LastTemp() (physic.Temperature, error) {
	// Read the scratchpad memory.
	spad, err := d.readScratchpad()
	if err != nil {
		return 0, err
	}

	// spad[1] is MSB, spad[0] is LSB and has 4 fractional bits. Need to do sign
	// extension multiply by 1000 to get Millis, divide by 16 due to 4 fractional
	// bits. Datasheet p.4.
	v := physic.Temperature((int(int8(spad[1]))<<8 + int(spad[0])))
	c := v*physic.Kelvin/16 + physic.ZeroCelsius

	// The device powers up with a value of 85°C, so if we read that odds are
	// very high that either no conversion was performed or that the conversion
	// failed due to lack of power. This prevents reading a temp of exactly 85°C,
	// but that seems like the right tradeoff.
	if c == 85000 {
		return 0, busError("ds18b20: has not performed a temperature conversion (insufficient pull-up?)")
	}

	return c, nil
}

//

// busError implements error and onewire.BusError.
type busError string

func (e busError) Error() string  { return string(e) }
func (e busError) BusError() bool { return true }

// conversionSleep sleeps for the time a conversion takes, which depends
// on the resolution:
// 9bits:94ms, 10bits:188ms, 11bits:376ms, 12bits:752ms, datasheet p.6.
func conversionSleep(bits int) {
	sleep((94 << uint(bits-9)) * time.Millisecond)
}

// readScratchpad reads the 9 bytes of scratchpad and checks the CRC.
// It returns the 8 bytes of scratchpad data (excluding the CRC byte).
func (d *Dev) readScratchpad() ([]byte, error) {
	// Read the scratchpad memory.
	var spad [9]byte
	if err := d.onewire.Tx([]byte{0xbe}, spad[:]); err != nil {
		return nil, err
	}

	// Check the scratchpad CRC.
	if !onewire.CheckCRC(spad[:]) {
		for _, s := range spad {
			if s != 0xff {
				return nil, busError("ds18b20: incorrect scratchpad CRC")
			}
		}
		return nil, busError("ds18b20: device did not respond")
	}

	return spad[:8], nil
}

var sleep = time.Sleep

var _ conn.Resource = &Dev{}
var _ physic.SenseEnv = &Dev{}

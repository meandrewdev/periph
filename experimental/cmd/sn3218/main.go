// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/meandrewdev/periph/conn/i2c/i2creg"
	"github.com/meandrewdev/periph/experimental/devices/sn3218"
	"github.com/meandrewdev/periph/host"
)

func main() {

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	d, err := sn3218.New(b)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Halt()

	if err := d.WakeUp(); err != nil {
		log.Fatal("Error while enabling device", err)
	}

	if err := d.BrightnessAll(1); err != nil {
		log.Fatal("Error while setting brightness", err)
	}

	// Switch LED 7 on
	if err := d.Switch(7, true); err != nil {
		log.Fatal("Error while switching LED", err)
	}
	time.Sleep(1000 * time.Millisecond)

	//Increase brightness for LED 7 to max
	if err := d.Brightness(7, 255); err != nil {
		log.Fatal("Error while changing LED brightness", err)
	}
	time.Sleep(1000 * time.Millisecond)

	//Get state of LED 7
	state, brightness, err := d.GetState(7)
	if err != nil {
		log.Fatal("Error while reading LED state", err)
	}
	log.Println("State: ", state, " - Brightness: ", brightness)

	// Switch all LEDs on
	if err := d.SwitchAll(true); err != nil {
		log.Fatal("Error while switching all LEDs", err)
	}
	time.Sleep(1000 * time.Millisecond)

	// Increase brightness for all
	if err := d.BrightnessAll(125); err != nil {
		log.Fatal("Error while changing globalBrightness", err)
	}
	time.Sleep(1000 * time.Millisecond)

	// Sleep mode to save energy, but keep state
	if err := d.Sleep(); err != nil {
		log.Fatal("Error while disabling device")
	}
	time.Sleep(1000 * time.Millisecond)

	// WakeUp again
	if err := d.WakeUp(); err != nil {
		log.Fatal("Error while enabling device")
	}
	time.Sleep(1000 * time.Millisecond)
}

// Copyright 2020 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package tlv493d_test

import (
	"fmt"
	"log"

	"github.com/meandrewdev/periph/conn/i2c/i2creg"
	"github.com/meandrewdev/periph/conn/physic"
	"github.com/meandrewdev/periph/experimental/devices/tlv493d"
	"github.com/meandrewdev/periph/host"
)

func Example() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open default I²C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("failed to open I²C: %v", err)
	}
	defer bus.Close()

	// Create a new TLV493D hall effect sensor.
	tlv, err := tlv493d.New(bus, &tlv493d.DefaultOpts)
	if err != nil {
		log.Fatalln(err)
	}
	defer tlv.Halt()

	// Read a single value.
	tlv.SetMode(tlv493d.LowPowerMode)
	fmt.Println("Single reading")
	reading, err := tlv.Read(tlv493d.HighPrecisionWithTemperature)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(reading)

	// Read values continuously from the sensor.
	fmt.Println("Continuous reading")
	c, err := tlv.ReadContinuous(100*physic.Hertz, tlv493d.LowPrecision)
	if err != nil {
		log.Fatalln(err)
	}

	for reading := range c {
		fmt.Println(reading)
	}
}

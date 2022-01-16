// Copyright 2019 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package inky_test

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/meandrewdev/periph/conn/gpio/gpioreg"
	"github.com/meandrewdev/periph/conn/spi/spireg"
	"github.com/meandrewdev/periph/experimental/devices/inky"
	"github.com/meandrewdev/periph/host"
)

func Example() {
	path := flag.String("image", "", "Path to image file (212x104) to display")
	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := spireg.Open("SPI0.0")
	if err != nil {
		log.Fatal(err)
	}

	dc := gpioreg.ByName("22")
	reset := gpioreg.ByName("27")
	busy := gpioreg.ByName("17")

	dev, err := inky.New(b, dc, reset, busy, &inky.Opts{
		Model:       inky.PHAT,
		ModelColor:  inky.Red,
		BorderColor: inky.Black,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := dev.Draw(img.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// spi-list lists all SPI ports.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/meandrewdev/periph/conn/pin"
	"github.com/meandrewdev/periph/conn/pin/pinreg"
	"github.com/meandrewdev/periph/conn/spi"
	"github.com/meandrewdev/periph/conn/spi/spireg"
)

func printPin(fn string, p pin.Pin) {
	name, pos := pinreg.Position(p)
	if name != "" {
		fmt.Printf("  %-4s: %-10s found on header %s, #%d\n", fn, p, name, pos)
	} else {
		fmt.Printf("  %-4s: %-10s\n", fn, p)
	}
}

func mainImpl() error {
	verbose := flag.Bool("v", false, "verbose mode")
	flag.Parse()
	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}
	log.SetFlags(log.Lmicroseconds)
	if flag.NArg() != 0 {
		return errors.New("unexpected argument, try -help")
	}

	if _, err := hostInit(); err != nil {
		return err
	}
	for _, ref := range spireg.All() {
		fmt.Printf("%s", ref.Name)
		if ref.Number != -1 {
			fmt.Printf(" #%d", ref.Number)
		}
		fmt.Print(":\n")
		if len(ref.Aliases) != 0 {
			fmt.Printf("  Aliases:\n")
			for _, a := range ref.Aliases {
				fmt.Printf("    %s\n", a)
			}
		}
		s, err := ref.Open()
		if err != nil {
			fmt.Printf("  Failed to open: %v\n", err)
			continue
		}
		if p, ok := s.(spi.Pins); ok {
			printPin("CLK", p.CLK())
			printPin("MOSI", p.MOSI())
			printPin("MISO", p.MISO())
			printPin("CS", p.CS())
		}
		if err := s.Close(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "spi-list: %s.\n", err)
		os.Exit(1)
	}
}

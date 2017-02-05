package main

import (
	"github.com/jtgoen/chip8-emu/emulator"
)

func main() {
	emulator := new(chip8emu.Chip8)

	emulator.Init()
	emulator.LoadApplication("/Users/jtgoen/golang/src/github.com/jtgoen/chip8-emu/runner/tetris.c8")

	for {
		emulator.EmulateCycle()

		if emulator.DrawFlag {

		}
	}
}

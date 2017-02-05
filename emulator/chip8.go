package chip8emu

import (
	"encoding/binary"
	"fmt"
	"os"
)

var chip8_fontset = [80]byte {
	0xF0, 0x90, 0x90, 0x90, 0xF0, //0
	0x20, 0x60, 0x20, 0x20, 0x70, //1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, //2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, //3
	0x90, 0x90, 0xF0, 0x10, 0x10, //4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, //5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, //6
	0xF0, 0x10, 0x20, 0x40, 0x40, //7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, //8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, //9
	0xF0, 0x90, 0xF0, 0x90, 0x90, //A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, //B
	0xF0, 0x80, 0x80, 0x80, 0xF0, //C
	0xE0, 0x90, 0x90, 0x90, 0xE0, //D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, //E
	0xF0, 0x80, 0xF0, 0x80, 0x80,  //F
}

type Chip8 struct {

	fontset [80]uint16

	DrawFlag bool

	Opcode uint16
	Memory [4096]byte

	V           [16]byte 		//CPU registers. 15 general-purpose, 1 'carry flag'
	I           uint16   		//Index Register
	PC          uint16   		//Program Counter
	GFX         [64 * 32]byte	//Represents graphics of the chip. 64 x 32 screen of white/black pixels.
	Delay_timer byte
	Sound_timer byte		//Tracks when the sound buzzer should trigger.

	Stack [16]uint16
	SP    uint16

	Key [16]byte
}

type Chip8_i interface {
	Init()
	EmulateCycle()
	DebugRender()
	LoadApplication() bool
}

func (c *Chip8) Init() {
	c.PC = 0x200
	c.Opcode = 0
	c.I = 0
	c.SP = 0

	// Clear display
	// Clear stack
	// Clear registers V0-VF
	// Clear memory

	// Load fontset
	for i := 0; i < 80; i++ {
		c.Memory[i] = chip8_fontset[i]
	}

	// Reset timers
}

func (c *Chip8) EmulateCycle() {
	c.Opcode = binary.BigEndian.Uint16(c.Memory[c.PC:c.PC+ 1])
}

func (c *Chip8) DebugRender() {

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c *Chip8) LoadApplication(filepath string) bool {
	c.Init()
	fmt.Printf("Loading: %s\n", filepath)

	file, err := os.Open(filepath)
	check(err)

	defer file.Close()

	fi, err := file.Stat()
	check(err)

	fmt.Printf("File size: %d\n", fi.Size())

	buffer := make([]byte, fi.Size())
	numRead, err := file.Read(buffer)
	check(err)

	fmt.Printf("Read %d bytes into buffer.", numRead)

	free_space := len(c.Memory) - 512

	if int64(free_space) > fi.Size() {
		for i := 0; i < len(buffer); i++ {
			c.Memory[i+512] = buffer[i]
		}
	} else {
		fmt.Println("Error: ROM too big for memory!")
	}

	return true
}
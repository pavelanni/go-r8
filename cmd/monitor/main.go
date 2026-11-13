package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pavelanni/go-r8"
	"golang.org/x/term"
)

func main() {
	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(fd, oldState)
	fmt.Print("Press [Enter] to step forward, or [Esc] to exit.\r\n")

	buf := make([]byte, 3)
	cpu := r8.NewCPU()
	// put opcode 48 (inc a) to the first three slots of Mem
	for i := range 3 {
		cpu.Mem[i] = 48
	}
	fmt.Printf("  PC   A")
	for {
		fmt.Printf("\n\r%04d %04d > ", cpu.PC, cpu.A)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if n == 1 && (buf[0] == 13 || buf[0] == 10) {
			cpu.Step()
		}
		if n == 1 && buf[0] == 27 {
			break
		}
	}
}

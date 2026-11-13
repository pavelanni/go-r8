// Package r8 emulates a simple CPU called the R8.
package r8

type CPU struct {
	PC  int
	A   int
	Mem [65536]int
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() {
	op := cpu.Mem[cpu.PC]
	switch op {
	case 48:
		cpu.A += 1
	default:
	}
	cpu.PC += 1
}

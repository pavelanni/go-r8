package r8_test

import (
	"reflect"
	"testing"

	"github.com/pavelanni/go-r8"
)

func TestNewInitialisesCPU(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	if cpu.PC != 0 {
		t.Errorf("after New, want pc == 0, got %d", cpu.PC)
	}
	got := cpu.Mem[0]
	if got != 0 {
		t.Errorf("after New, want Mem[0] == 0, got %d", got)
	}
}

func TestStepIncrementsPC(t *testing.T) {
	t.Parallel()
	cpu := r8.NewCPU()
	cpu.Step()
	if cpu.PC != 1 {
		t.Errorf("want pc == 1, got %d", cpu.PC)
	}
	cpu.Step()
	if cpu.PC != 2 {
		t.Errorf("want pc == 2, got %d", cpu.PC)
	}
}

func TestRegisterAExistsAndZero(t *testing.T) {
	rt := reflect.TypeFor[r8.CPU]()
	_, found := rt.FieldByName("A")
	if !found {
		t.Errorf("want A in CPU struct, got nothing")
		return
	}
	cpu := r8.NewCPU()
	val := reflect.ValueOf(cpu)
	regA := val.Elem().FieldByName("A")
	switch regA.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	// do nothing
	default:
		t.Errorf("want A to be Int, got %v", regA.Kind())
		return
	}
	regAint := regA.Int()
	if regAint != 0 {
		t.Errorf("want A == 0, got %d", regAint)
	}
}

func TestInstructionInc(t *testing.T) {
	cpu := r8.NewCPU()
	cpu.Mem[0] = 48
	cpu.Step()
	if cpu.A != 1 {
		t.Errorf("want A == 1, got %d", cpu.A)
	}
}

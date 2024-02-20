package main

import (
	"github.com/YujiNNakashima/state-machines/machines"
)

func main() {

	// sm := machines.InstantiateSemaphoreStateMachine(&machines.RedLight{})
	sm := machines.InstantiateVendingMachine()

	for {
		sm.Transition()
	}
}

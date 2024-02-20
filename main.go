package main

import (
	"github.com/YujiNNakashima/state-machines/machines"
)

func main() {

	sm := machines.InstantiateSemaphoreStateMachine(&machines.RedLight{})

	for {
		sm.Transition()
	}
}

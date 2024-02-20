package machines

import (
	"fmt"
	"time"
)

type SemaphoreStateMachine struct {
	currentState State
	states       map[string]State
}

func (sm *SemaphoreStateMachine) setState(s State) {
	sm.currentState = s
	sm.currentState.Enter()
}

func (sm *SemaphoreStateMachine) Transition() {
	sm.currentState.Update(sm)
}

type State interface {
	Enter()
	Exit()
	Update(sm *SemaphoreStateMachine)
}

type RedLight struct{}

func (s *RedLight) Enter() {
	fmt.Println("RED LIGHT state, stop!")
	time.Sleep(time.Second * 5)
}

func (s *RedLight) Exit() {
}

func (s *RedLight) Update(sm *SemaphoreStateMachine) {
	fmt.Println("updating red to green...")
	sm.setState(&GreenLight{})
}

type GreenLight struct{}

func (s *GreenLight) Enter() {
	fmt.Println("GREEN LIGHT state, move your ass...")
	time.Sleep(time.Second * 5)
}

func (s *GreenLight) Exit() {
}

func (s *GreenLight) Update(sm *SemaphoreStateMachine) {
	fmt.Println("updating green to yellow...")
	sm.setState(&YellowLight{})
}

type YellowLight struct{}

func (s *YellowLight) Enter() {
	fmt.Println("YELLOW LIGHT state, wait...")
	time.Sleep(time.Second * 2)
}

func (s *YellowLight) Exit() {
}

func (s *YellowLight) Update(sm *SemaphoreStateMachine) {
	fmt.Println("updating yellow to green...")
	sm.setState(&RedLight{})
}

func InstantiateSemaphoreStateMachine(initialState State) *SemaphoreStateMachine {
	sm := &SemaphoreStateMachine{
		currentState: initialState,
		states:       make(map[string]State),
	}

	sm.currentState.Enter()
	return sm
}

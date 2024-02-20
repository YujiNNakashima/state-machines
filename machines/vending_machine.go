package machines

import (
	"fmt"
	"time"
)

type VendingMachine struct {
	currentState VendingState
	states       map[string]VendingState
}

func (vm *VendingMachine) setState(s VendingState) {
	vm.currentState = s
	vm.currentState.Enter()
}

func (vm *VendingMachine) Transition() {
	vm.currentState.Update(vm)
}

type VendingState interface {
	Enter()
	Exit()
	Update(vm *VendingMachine)
}

type IdleState struct{}

func (s *IdleState) Enter() {
	fmt.Println("Vending machine is idle.")
}

func (s *IdleState) Exit() {}

func (s *IdleState) Update(vm *VendingMachine) {
	// Handle user input to select an item
	fmt.Println("Please select an item (1-3):")
	var item int
	fmt.Scanln(&item)
	switch item {
	case 1, 2, 3:
		vm.setState(&SelectingItemState{Item: item})
	default:
		fmt.Println("Invalid item selection.")
	}
}

type SelectingItemState struct {
	Item int
}

func (s *SelectingItemState) Enter() {
	fmt.Printf("Selected item: %d\n", s.Item)
	fmt.Println("Processing item selection...")
}

func (s *SelectingItemState) Exit() {}

func (s *SelectingItemState) Update(vm *VendingMachine) {
	// Simulate processing the item selection
	time.Sleep(time.Second * 2)
	vm.setState(&DispensingItemState{})
}

type DispensingItemState struct{}

func (s *DispensingItemState) Enter() {
	fmt.Println("Dispensing item...")
}

func (s *DispensingItemState) Exit() {}

func (s *DispensingItemState) Update(vm *VendingMachine) {
	// Simulate dispensing the item
	time.Sleep(time.Second * 2)
	fmt.Println("Item dispensed successfully.")
	vm.setState(&IdleState{})
}

func InstantiateVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		currentState: &IdleState{},
		states:       make(map[string]VendingState),
	}

	vm.currentState.Enter()
	return vm
}

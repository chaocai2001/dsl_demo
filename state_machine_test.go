package statemachine

import (
	"fmt"
	"testing"
)

func TestStateMachine(t *testing.T) {
	// Build an Auto-Door
	autoDoor := StateMachine("auto-door").
		States(
			State("Open").Transitions(
				Transtion().TriggedBy("timeout").ToState("Close"),
			),
			State("Close").Transitions(
				Transtion().TriggedBy("detected").ToState("Open"),
			),
		)

	// Test the Auto-Door
	fmt.Print("Send 'timeout' event \t")
	autoDoor.ReceivedEvent("timeout")
	fmt.Println("current state:", autoDoor.curState.Name())
	fmt.Print("Send 'detected' event \t")
	autoDoor.ReceivedEvent("detected")
	fmt.Println("current state:", autoDoor.curState.Name())
	fmt.Print("Send 'timeout' event \t")
	autoDoor.ReceivedEvent("timeout")
	fmt.Println("current state:", autoDoor.curState.Name())
}

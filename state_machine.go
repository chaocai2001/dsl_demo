package statemachine

type EventType string

func Event(evt string) EventType {
	return EventType(evt)
}

func TriggedBy(evt string) EventType {
	return EventType(evt)
}

type TransitionST struct {
	triggedBy EventType
	toState   string
}

func Transtion() *TransitionST {
	return &TransitionST{}
}

func (tran *TransitionST) TriggedBy(evt string) *TransitionST {
	tran.triggedBy = Event(evt)
	return tran
}

func (tran *TransitionST) ToState(state string) *TransitionST {
	tran.toState = state
	return tran
}

type StateST struct {
	name        string
	transitions []*TransitionST
}

func State(name string) *StateST {
	return &StateST{
		name:        name,
		transitions: []*TransitionST{},
	}
}

func (state *StateST) Transitions(otherTrans ...*TransitionST) *StateST {
	//state.transitions = append(state.transitions, trans)
	for _, tran := range otherTrans {
		state.transitions = append(state.transitions, tran)
	}
	return state
}

func (state *StateST) Name() string {
	return state.name
}

type StateMachineST struct {
	name     string
	states   []*StateST
	curState *StateST
}

func StateMachine(name string) *StateMachineST {
	return &StateMachineST{
		name: name,
	}
}

func (sm *StateMachineST) States(states ...*StateST) *StateMachineST {
	sm.states = states
	sm.curState = states[0]
	return sm
}

func (sm *StateMachineST) ReceivedEvent(evt EventType) {
	for _, tran := range sm.curState.transitions {
		if tran.triggedBy == evt {
			nextStateName := tran.toState
			for _, state := range sm.states {
				if state.name == nextStateName {
					sm.curState = state
					return
				}

			}
			panic("invald state: " + nextStateName)
		}
	}
	panic("invald event: " + evt)
}

func (sm *StateMachineST) CurrentState() *StateST {
	return sm.curState
}

// Code generated by reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

type AppElem struct {
	react.Element
}

func buildApp(cd react.ComponentDef) react.Component {
	return AppDef{ComponentDef: cd}
}

func buildAppElem(children ...react.Element) *AppElem {
	return &AppElem{
		Element: react.CreateElement(buildApp, nil, children...),
	}
}

// SetState is an auto-generated proxy proxy to update the state for the
// App component.  SetState does not immediately mutate a.State()
// but creates a pending state transition.
func (a AppDef) SetState(state AppState) {
	a.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the App component
func (a AppDef) State() AppState {
	return a.ComponentDef.State().(AppState)
}

// IsState is an auto-generated definition so that AppState implements
// the myitcv.io/react.State interface.
func (a AppState) IsState() {}

var _ react.State = AppState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (a AppDef) GetInitialStateIntf() react.State {
	return AppState{}
}

func (a AppState) EqualsIntf(val react.State) bool {
	return a == val.(AppState)
}

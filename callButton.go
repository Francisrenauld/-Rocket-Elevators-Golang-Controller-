package main

//Button on a floor or basement to go back to lobby
type CallButton struct {

	ID int
    floor int
    direction string
}

func NewCallButton(callButtonID, _floor int, _direction string) *CallButton {

	btn := &CallButton{

		ID: callButtonID,
		floor: _floor,
		direction: _direction,

	}

return btn
}

package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {

	ID int;
	floor int;
	status string;
	direction string;

}

func NewFloorRequestButton(_id, _floor int, _direction string) *FloorRequestButton{

	f := &FloorRequestButton{

		ID: _id,
		floor: _floor,
		direction: _direction,

	}
return f
}

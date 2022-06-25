package main

type Door struct {

	ID string
	status string
}

func NewDoor(_id, _status string) *Door {

	d := &Door{

		ID: _id,
		status: _status,
	}
return d
}

package main

import "sort"


type BestElevatorInformations struct{

bestElevator Elevator
bestScore int
referenceGap int

}

func NewBestElevatorInformations(_bestElevator Elevator, _bestScore, _referenceGap int) *BestElevatorInformations {

return &BestElevatorInformations{_bestElevator, _bestScore, _referenceGap}


}
type Elevator struct {

	ID string
	status string
	amountOfFloors int
	direction string
	currentFloor int
	overweight bool
	completedRequestsList []int
	floorRequestsList []int
	
}

func NewElevator(_elevatorID, _status string,  _amountOfFloors,  _currentFloor int) *Elevator {

	e := &Elevator{

		ID: _elevatorID,
		status: _status,
		amountOfFloors: _amountOfFloors,
		direction: "nil",
		currentFloor: _currentFloor,
		overweight: false,
	}
	return e
}

func (e *Elevator) move() {
	
	var destination int 
	
	for len(e.floorRequestsList) > 0{

		destination = e.floorRequestsList[0]
		e.status = "moving"

		if e.currentFloor < destination{

			e.direction = "up"
			e.sortFloorList()
			destination = e.floorRequestsList[0]

			for e.currentFloor < destination{

				e.currentFloor++
			}
		}else if e.currentFloor > destination{

			e.direction = "down";
			e.sortFloorList();
			destination = e.floorRequestsList[0];

			for e.currentFloor > destination{
				e.currentFloor--;
			}
		}
		e.status = "stopped";
		//operateDoors();
		e.completedRequestsList = append(e.completedRequestsList, destination)
		remove(e.floorRequestsList, 0)


	}
}

func (e *Elevator) sortFloorList(){

	if e.direction == "up"{

		sort.Ints(e.floorRequestsList)
	}else{

		sort.Ints(e.floorRequestsList)
		reverse(e.floorRequestsList)

	}
}

func (e *Elevator) addNewRequest(requestedFloor int){

	if contains(e.floorRequestsList, requestedFloor) == false{

		e.floorRequestsList = append(e.floorRequestsList,requestedFloor)

	}
	if e.currentFloor < requestedFloor{

		e.direction = "up"

	}
	if e.currentFloor > requestedFloor{

		e.direction = "down"

	}

}

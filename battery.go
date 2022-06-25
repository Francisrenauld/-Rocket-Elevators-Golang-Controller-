package main

import (
	"math"
	"sort"
)

var floor = 1 
var columnID = 1

type Battery struct {
	
	ID int
	floorButtonsList []FloorRequestButton
	columnsList []Column
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery{
	
	b := Battery{
		ID:      _id,
		floorButtonsList: []FloorRequestButton{},
		columnsList: []Column{},
	}


	if _amountOfBasements > 0 {

		b.CreateBasementFloorRequestButtons(_amountOfBasements)
		b.CreateBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn);
		_amountOfColumns--
	}

	b.CreateFloorRequestButtons(_amountOfFloors);
	b.CreateColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn);

	
	return &b
}

func (b *Battery)CreateBasementColumn( _amountOfBasements, _amountOfElevatorPerColumn int){

	var servedFloors []int

	floor = -1
	servedFloors = append(servedFloors, 1)
	for  i  := 0; i < _amountOfBasements; i++{

		servedFloors = append(servedFloors, floor)
		floor--;

	}
	
	column := NewColumn(columnID, _amountOfBasements, _amountOfElevatorPerColumn, servedFloors, true)
	b.columnsList = append(b.columnsList, *column)
	columnID++;
}

func (b *Battery)CreateColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn int){

	var moean float64 = float64(_amountOfFloors / _amountOfColumns);
	var amountOfFloorsPerColumn = math.Ceil(moean)
	floor = 1

	for i := 0; i < _amountOfColumns; i++{
		var counter = 0
		var servedFloors []int
		//
		for j := 0; float64(j) < amountOfFloorsPerColumn; j++{
			
			if floor <= _amountOfFloors{

				servedFloors = append(servedFloors, floor)
				floor++;
				counter++
			}
		}
		if !contains(servedFloors, 1){
			servedFloors = append(servedFloors, 1)
			sort.Ints(servedFloors)
		}
		column := NewColumn(columnID, counter, _amountOfElevatorPerColumn, servedFloors, false)
		b.columnsList = append(b.columnsList, *column)
		columnID++;
	}
}

func (b *Battery) CreateFloorRequestButtons (_amountOfFloors int){

	buttonFloor := 1
	floorRequestButtonID := 0;

	for i := 0; i < _amountOfFloors; i++{

		floorRequestButtonID = i + 1
		
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "up")
		b.floorButtonsList = append(b.floorButtonsList, *floorRequestButton)

		buttonFloor++
	}
}

func (b *Battery) CreateBasementFloorRequestButtons (_amountOfBasements int){

	 buttonFloor := -1;
	 floorRequestButtonID := 0;

	 for  i := 0; i < _amountOfBasements; i++{

		floorRequestButtonID = i +1
		
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "down")
		b.floorButtonsList = append(b.floorButtonsList, *floorRequestButton)

		buttonFloor--
	 }
}



func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	
	var bestColumn Column

	for  _, column := range b.columnsList{

		if contains(column.servedFloorsList, _requestedFloor) == true{

			bestColumn = column
		}
	}
	return &bestColumn
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {

	column := *b.findBestColumn(_requestedFloor)
	elevator := *column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
    elevator.addNewRequest(_requestedFloor)
    elevator.move()

	return &column, &elevator
}



package main

var buttonFloor = 0
//var callButtonID = 1
var elevatorID = 1
  
type Column struct {

	ID int
	status string 
	amountOfElevators int
	servedFloorsList []int
	//amountOfFLoors int
	elevatorsList []*Elevator
	callButtonsList []*CallButton
	

}

func NewColumn(_id, _amountOfFloors, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column{

	c := &Column{

		ID:		      _id,
		status:  "online",
		amountOfElevators: _amountOfElevators,
		servedFloorsList: _servedFloors,

	}
	c.CreateElevators(_amountOfFloors, _amountOfElevators)
	c.CreateCallButtons(_amountOfFloors, _isBasement)
	return c
}

func (c *Column) CreateCallButtons(_amountOfFloors int,  _isBasement bool){

	var callButtonID = 1
	if _isBasement == true{

		buttonFloor = -1
		for	i  := 0; i < _amountOfFloors; i++{

			callButton := NewCallButton(callButtonID, buttonFloor, "down")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor--
			callButtonID++
		}
	}else{

		buttonFloor = 1
		
		for	i  := 0; i < _amountOfFloors; i++{

			callButton := NewCallButton(callButtonID, buttonFloor, "up")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor++
			callButtonID++
		}
	}
}

func (c *Column) CreateElevators( _amountOfFloors, _amountOfElevators int){
	var elevatorID = 1
	for	i  := 0; i < _amountOfElevators; i++{

		elevator := NewElevator(toCharStrArr(elevatorID), "idle", _amountOfFloors, 1)
		c.elevatorsList = append(c.elevatorsList, elevator)

		elevatorID++
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(userPosition int, direction string) *Elevator {

	var elevator Elevator
	elevator = *c.findElevator(userPosition, direction)
	elevator.addNewRequest(userPosition)
    elevator.move()
    elevator.addNewRequest(1)
    elevator.move()

	return &elevator
}


func (c *Column) findElevator ( requestedFloor int,  requestedDirection string) *Elevator {

	var bestElevator Elevator = *c.elevatorsList[0]
	var bestScore = 5
	var referenceGap = 10000000
	var bestElevatorInfo BestElevatorInformations

	if requestedFloor < 0{

		for  _, elevator := range c.elevatorsList{

			if 1 == elevator.currentFloor && elevator.status == "stopped"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(1, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else if 1 == elevator.currentFloor && elevator.status == "idle"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(2, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else if 1 > elevator.currentFloor && elevator.direction == "up"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(2, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else if 1 < elevator.currentFloor && elevator.direction == "down"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(3, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)
				
			}else if elevator.status == "idle"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(4, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(5, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}
		}
		bestElevator = bestElevatorInfo.bestElevator
		bestScore = bestElevatorInfo.bestScore
		referenceGap = bestElevatorInfo.referenceGap
	}else{

		for _, elevator := range c.elevatorsList{

			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(1, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else if requestedFloor > elevator.currentFloor && elevator.status == "stopped" && elevator.direction == "up" && requestedDirection == "up"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(2, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(2, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else if elevator.status == "idle"{

				bestElevatorInfo = *c.checkIfElevatorIsBetter(4, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)

			}else if elevator.direction != requestedDirection{
				
				bestElevatorInfo = *c.checkIfElevatorIsBetter(5, *elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}

			bestElevator = bestElevatorInfo.bestElevator
			bestScore = bestElevatorInfo.bestScore
			referenceGap = bestElevatorInfo.referenceGap
		}
	}
	return &bestElevator
}

func (c *Column) checkIfElevatorIsBetter( scoreToCheck int,  newElevator Elevator,  bestScore,  referenceGap int,  bestElevator Elevator,  floor int) *BestElevatorInformations{

var gap = 0

	if scoreToCheck < bestScore{

		bestScore = scoreToCheck
		bestElevator = newElevator
		var refGap = (newElevator.currentFloor - floor)
		referenceGap = Abs(refGap)
	}else if bestScore == scoreToCheck{

		var gap1 = (newElevator.currentFloor - floor)
		gap = Abs(gap1)

		if referenceGap > gap{

			bestElevator = newElevator
            referenceGap = gap

		}
	}
	return NewBestElevatorInformations(bestElevator, bestScore, referenceGap)
}

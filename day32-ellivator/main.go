package main

import (
	"fmt"
)

type Ellivator interface {
	PickUp() error
	Drop() error
	CountOfPeoplesInBuilding() (int, error)
}

type ellivator struct {
	CapacityOfPersons            int32
	NoOfFlowers                  int32
	BuildingName                 string
	CountOfNoPeopleUsedEllivator int32
	CountOfPeopleInBuilding      int32
	AvgFamilyMembersCountPerFlat int32
}

func (e *ellivator) PickUp() error {

	return nil
}

func (e *ellivator) Drop() error {
	return nil
}

func (e *ellivator) CountOfPeoplesInBuilding(totalFlower int32,flowerWiseNoOfFlats int32) (int, error) {
	count := (flowerWiseNoOfFlats*totalFlower) * 3
	return count, nil
}

func main() {
	fmt.Println("Welcome to Serenity Apartment...!")
	var (
		FlowerWiseNoOfFlats          int32 = 7
		TotalFlower                  int32 = 4
		TotalFlats                   int32 = FlowerWiseNoOfFlats * TotalFlower
		AvgFamilyMembersCountPerFlat int32 = 3 * TotalFlats
	)
	countOfPeoplesInBuilding, err := CountOfPeoplesInBuilding(TotalFlower,FlowerWiseNoOfFlats)
	if err != nil{
		err = errorf("error while fetching countOfPeoplesInBuilding - ", err)
		return 0, nil
	}
	fmt.Println("FlowerWiseNoOfFlats= ", 7)
	fmt.Println("TotalFlower= ", 4)
	fmt.Println("TotalFlats= ", 28)
	fmt.Println("AvgFamilyMemberCountPerFlat= ", AvgFamilyMembersCountPerFlat)

	fmt.Println("CountOfPeoplesInBuilding = ",countOfPeoplesInBuilding)
	elv := &ellivator{
		CapacityOfPersons:            6,
		NoOfFlowers:                  4,
		BuildingName:                 "Serenity",
		CountOfNoPeopleUsedEllivator: 56, // mostly daily use of ellivator is made by mother and kids
		CountOfPeopleInBuilding:      84,
		AvgFamilyMembersCountPerFlat: 3 * TotalFlats,
	}
	fmt.Println("ellivator = ", elv)
}

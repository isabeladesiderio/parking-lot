package internal_test

import (
	"testing"

	internal "github.com/isabeladesiderio/internal/parkinglot"
)

func TestIsFull(t *testing.T) {

	t.Run("1 - Is full", func(t *testing.T) {
		parkingLot := internal.ParkingLot{TotalSpaces: 0}

		full := parkingLot.IsFull()

		if !full {
			t.Errorf("have: %#v, want: %#v", full, true)
		}
	})

	t.Run("2 - Isn't full", func(t *testing.T) {
		parkingLot := internal.ParkingLot{TotalSpaces: 10}

		full := parkingLot.IsFull()

		if full {
			t.Errorf("have: %#v, want: %#v", full, false)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("1 - Is empty", func(t *testing.T) {
		parkingLot := internal.ParkingLot{TotalSpaces: 5}
		empty := parkingLot.IsEmpty()

		if !empty {
			t.Errorf("have: %#v, want: %#v", empty, true)
		}
	})

	t.Run("2 - Isn't empty", func(t *testing.T) {
		parkingLot := internal.ParkingLot{TotalSpaces: 5, OccupiedSpaces: 3}
		empty := parkingLot.IsEmpty()

		if empty {
			t.Errorf("have: %#v, want: %#v", empty, false)
		}
	})
}

func TestGetSpacesByVehicle(t *testing.T) {

	t.Run("1 - Spaces available for cars", func(t *testing.T) {
		parkingLot := internal.ParkingLot{
			TotalSpaces:  10,
			VehiclesType: map[internal.VehicleType]int{internal.Car: 5, internal.Motorcycle: 2, internal.Van: 3},
		}
		availableSpaces := parkingLot.GetSpacesByVehicle(internal.Car)
		expectedAvailableSpaces := 5

		if availableSpaces != expectedAvailableSpaces {
			t.Errorf("have: %#v, want: %#v", availableSpaces, expectedAvailableSpaces)
		}
	})

	t.Run("2 - Spaces available for vans", func(t *testing.T) {
		parkingLot := internal.ParkingLot{
			TotalSpaces:  10,
			VehiclesType: map[internal.VehicleType]int{internal.Car: 5, internal.Motorcycle: 2, internal.Van: 3},
		}
		availableSpaces := parkingLot.GetSpacesByVehicle(internal.Van)
		expectedAvailableSpaces := 3

		if availableSpaces != expectedAvailableSpaces {
			t.Errorf("have: %#v, want: %#v", availableSpaces, expectedAvailableSpaces)
		}
	})

	t.Run("3 - Spaces available for moto", func(t *testing.T) {
		parkingLot := internal.ParkingLot{
			TotalSpaces:  10,
			VehiclesType: map[internal.VehicleType]int{internal.Car: 5, internal.Motorcycle: 2, internal.Van: 3},
		}
		availableSpaces := parkingLot.GetSpacesByVehicle(internal.Motorcycle)
		expectedAvailableSpaces := 2

		if availableSpaces != expectedAvailableSpaces {
			t.Errorf("have: %#v, want: %#v", availableSpaces, expectedAvailableSpaces)
		}
	})

}

func TestSpacesVehicleTypeFull(t *testing.T) {

	t.Run("1 - Spaces are not full for the specified vehicle type", func(t *testing.T) {
		parkingLot := internal.ParkingLot{
			VehiclesType: map[internal.VehicleType]int{internal.Car: 2, internal.Motorcycle: 1, internal.Van: 0},
		}
		full := parkingLot.SpacesVehicleTypeFull(internal.VehicleType(internal.Van))

		if !full {
			t.Errorf("have: %#v, want: %#v", full, false)
		}
	})

	t.Run("2 - Spaces are full for the specified vehicle type", func(t *testing.T) {
		parkingLot := internal.ParkingLot{
			VehiclesType: map[internal.VehicleType]int{internal.Car: 0, internal.Motorcycle: 0, internal.Van: 0},
		}
		full := parkingLot.SpacesVehicleTypeFull(internal.VehicleType(internal.Car))

		if !full {
			t.Errorf("have: %#v, want: %#v", full, true)
		}
	})
}

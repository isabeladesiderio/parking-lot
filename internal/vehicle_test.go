package internal_test

import (
	"testing"

	"github.com/isabeladesiderio/internal"
)

func TestDistributeInitialSpaces(t *testing.T) {

	var testScenarios = []struct {
		Name         string
		TotalSpaces  int
		expected     map[internal.VehicleType]int
		expectedCar  int
		expectedMoto int
		expectedVan  int
	}{
		{
			Name:        "1 - Total spaces",
			TotalSpaces: 8,
			expected: map[internal.VehicleType]int{
				internal.Car:        4,
				internal.Motorcycle: 2,
				internal.Van:        2,
			},
			expectedCar:  4,
			expectedMoto: 2,
			expectedVan:  2,
		},
		{
			Name:        "2 - Total spaces",
			TotalSpaces: 16,
			expected: map[internal.VehicleType]int{
				internal.Car:        8,
				internal.Motorcycle: 4,
				internal.Van:        4,
			},
			expectedCar:  8,
			expectedMoto: 4,
			expectedVan:  4,
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.Name, func(t *testing.T) {
			vehicleTypes := internal.DistributeInitialSpaces(ts.TotalSpaces)

			if vehicleTypes[internal.Car] != ts.expectedCar ||
				vehicleTypes[internal.Motorcycle] != ts.expectedMoto ||
				vehicleTypes[internal.Van] != ts.expectedVan {
				t.Errorf("have: %#v, %#v, %#v, want: %#v, %#v, %#v",
					vehicleTypes[internal.Car], vehicleTypes[internal.Motorcycle], vehicleTypes[internal.Van],
					ts.expectedCar, ts.expectedMoto, ts.expectedVan)
			}
		})
	}
}

func TestGetOccupiedSpacesByVehicle(t *testing.T) {

	var testScenarios = []struct {
		Name         string
		TotalSpaces  int
		VehicleType  internal.VehicleType
		SpacesByType map[internal.VehicleType]int
		expected     int
	}{
		{
			Name:         "1 - Car spaces",
			TotalSpaces:  10,
			VehicleType:  internal.Car,
			SpacesByType: map[internal.VehicleType]int{internal.Car: 4, internal.Motorcycle: 2, internal.Van: 2},
			expected:     1,
		},
		{
			Name:         "2 - Car spaces ",
			TotalSpaces:  10,
			VehicleType:  internal.Car,
			SpacesByType: map[internal.VehicleType]int{internal.Car: 3, internal.Motorcycle: 2, internal.Van: 2},
			expected:     2,
		},
		{
			Name:         "3 - Motorcycle spaces",
			TotalSpaces:  8,
			VehicleType:  internal.Motorcycle,
			SpacesByType: map[internal.VehicleType]int{internal.Car: 4, internal.Motorcycle: 2, internal.Van: 2},
			expected:     0,
		},
		{
			Name:         "4 - Motorcycle spaces",
			TotalSpaces:  8,
			VehicleType:  internal.Motorcycle,
			SpacesByType: map[internal.VehicleType]int{internal.Car: 4, internal.Motorcycle: 0, internal.Van: 2},
			expected:     2,
		},
		{
			Name:         "5 - Van spaces",
			TotalSpaces:  8,
			VehicleType:  internal.Van,
			SpacesByType: map[internal.VehicleType]int{internal.Car: 4, internal.Motorcycle: 2, internal.Van: 2},
			expected:     0,
		},
		{
			Name:         "6 - Van spaces",
			TotalSpaces:  8,
			VehicleType:  internal.Van,
			SpacesByType: map[internal.VehicleType]int{internal.Car: 4, internal.Motorcycle: 2, internal.Van: 1},
			expected:     1,
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.Name, func(t *testing.T) {
			parkingLot := &internal.ParkingLot{
				TotalSpaces:  ts.TotalSpaces,
				VehiclesType: ts.SpacesByType,
			}

			occupiedSpaces := parkingLot.GetOccupiedSpacesByVehicle(ts.VehicleType)

			if occupiedSpaces != ts.expected {
				t.Errorf("have: %#v, want: %#v", occupiedSpaces, ts.expected)
			}
		})
	}
}

package internal

import "fmt"

func DistributeSpaces(TotalSpaces int) map[VehicleType]int {
	vehiclesType := make(map[VehicleType]int)
	vehiclesType[Car] = TotalSpaces / 2
	vehiclesType[Motorcycle] = TotalSpaces / 4
	vehiclesType[Van] = TotalSpaces / 4

	return vehiclesType
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) error {

	// Check if the parking lot is full
	if p.IsFull() {
		return fmt.Errorf("parking lot is full")
	}

	// Checks if there is already a vehicle with the same license plate in the parking lot
	for _, existsVehicle := range p.VehiclesDescription {
		if existsVehicle.LicensePlate == vehicle.LicensePlate {
			return fmt.Errorf("vehicle with the same license plate is already parked")
		}
	}

	if vehicle.Type == Van {
		// Check if there are available van spaces
		if p.VehiclesType[Van] > 0 {
			p.VehiclesType[Van]--
		} else if p.VehiclesType[Car] >= 3 {
			// If there is no available van space, park in 3 car spaces
			p.VehiclesType[Car] -= 3
		} else {
			return fmt.Errorf("no available spaces for this vehicle type")
		}
	}

	if vehicle.Type == Car {
		// Check if there are available car spaces
		if p.VehiclesType[Car] > 0 {
			p.VehiclesType[Car]--
		} else {
			return fmt.Errorf("no available spaces for cars")
		}
	}
	if vehicle.Type == Motorcycle {
		// Check if there are available motorcycle spaces
		if p.VehiclesType[Motorcycle] > 0 {
			p.VehiclesType[Motorcycle]--
		} else {
			return fmt.Errorf("no available spaces for moto")
		}
	}

	// Register the vehicle in the vehicle description map
	p.VehiclesDescription[vehicle.LicensePlate] = vehicle
	p.OccupiedSpaces++

	return nil
}

func (p *ParkingLot) GetOccupiedSpacesByVehicle(vehicleType VehicleType) int {
	initialSpaces := p.TotalSpaces / 2
	if vehicleType == Motorcycle || vehicleType == Van {
		initialSpaces = p.TotalSpaces / 4
	}

	occupiedSpaces := initialSpaces - p.GetSpacesByVehicle(vehicleType)
	return occupiedSpaces
}

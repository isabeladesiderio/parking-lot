package main

import (
	internal "github.com/isabeladesiderio/internal/parkinglot"
)

func main() {

	// Example usage of Parking Lot
	parkingLot := internal.NewParkingLot(10)

	vehicle1 := internal.Vehicle{LicensePlate: "CAR123", Type: internal.Car}
	vehicle2 := internal.Vehicle{LicensePlate: "MOTO456", Type: internal.Motorcycle}
	vehicle3 := internal.Vehicle{LicensePlate: "VAN789", Type: internal.Van}

	parkingLot.ParkVehicle(vehicle1)
	parkingLot.ParkVehicle(vehicle2)
	parkingLot.ParkVehicle(vehicle3)

	internal.ParkingLotStatus(parkingLot)

}

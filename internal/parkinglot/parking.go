package internal

import "fmt"

type VehicleType int

const (
	Car        VehicleType = 0
	Motorcycle VehicleType = 1
	Van        VehicleType = 2
)

// Structure for the vehicles
type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

// Structure for the parking lot
type ParkingLot struct {
	TotalSpaces         int
	OccupiedSpaces      int
	VehiclesType        map[VehicleType]int
	VehiclesDescription map[string]Vehicle
}

// Function for formatting and printing parking information
func formatInfo(label string, value interface{}) {
	fmt.Printf("%s: %v\n", label, value)
}

func ParkingLotStatus(parkingLot *ParkingLot) {
	// Vehicle Spaces
	formatInfo("Empty vehicle spaces", parkingLot.TotalSpaces-parkingLot.OccupiedSpaces)
	formatInfo("Total vehicle spaces", parkingLot.TotalSpaces)

	// Parking lot
	formatInfo("Is parking lot full?", parkingLot.IsFull())
	formatInfo("Is parking lot empty?", parkingLot.IsEmpty())

	// Veiculos spaces full
	formatInfo("Are motorcycle spaces full?", parkingLot.SpacesVehicleTypeFull(Motorcycle))
	formatInfo("Are car spaces full?", parkingLot.SpacesVehicleTypeFull(Car))
	formatInfo("Are van spaces full?", parkingLot.SpacesVehicleTypeFull(Van))

	// Available spaces
	formatInfo("Available spaces for cars", parkingLot.GetSpacesByVehicle(Car))
	formatInfo("Available spaces for vans", parkingLot.GetSpacesByVehicle(Van))
	formatInfo("Available spaces for motorcycles", parkingLot.GetSpacesByVehicle(Motorcycle))

	// Occupied spaces
	formatInfo("Occupied spaces for cars", parkingLot.GetOccupiedSpacesByVehicle(Car))
	formatInfo("Occupied spaces for vans", parkingLot.GetOccupiedSpacesByVehicle(Van))
	formatInfo("Occupied spaces for motorcycles", parkingLot.GetOccupiedSpacesByVehicle(Motorcycle))
}

func NewParkingLot(totalSpaces int) *ParkingLot {
	prk := ParkingLot{
		TotalSpaces:         totalSpaces,
		OccupiedSpaces:      0,
		VehiclesType:        DistributeSpaces(totalSpaces),
		VehiclesDescription: make(map[string]Vehicle),
	}

	return &prk
}

func (p *ParkingLot) IsFull() bool {
	return p.OccupiedSpaces >= p.TotalSpaces
}

func (p *ParkingLot) IsEmpty() bool {
	return p.OccupiedSpaces == 0
}

func (p *ParkingLot) SpacesVehicleTypeFull(vehicleType VehicleType) bool {
	return p.VehiclesType[vehicleType] == 0
}

func (p *ParkingLot) GetSpacesByVehicle(vehicleType VehicleType) int {
	return p.VehiclesType[vehicleType]
}

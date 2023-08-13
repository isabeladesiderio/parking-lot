package internal

import "fmt"

type VehicleType int

const (
	Car        VehicleType = 0
	Motorcycle VehicleType = 1
	Van        VehicleType = 2
)

// Definindo a estrutura para um veículo
type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

// Definindo a estrutura para o estacionamento
type ParkingLot struct {
	TotalSpaces         int
	OccupiedSpaces      int
	VehiclesType        map[VehicleType]int
	VehiclesDescription map[string]Vehicle
}

func ParkingLotStatus(parkingLot *ParkingLot) {
	fmt.Printf("Empty spaces: %d\n", parkingLot.TotalSpaces-parkingLot.OccupiedSpaces)
	fmt.Printf("Total spaces: %d\n", parkingLot.TotalSpaces)
	fmt.Println("Is parking lot full?", parkingLot.IsFull())
	fmt.Println("Is parking lot empty?", parkingLot.IsEmpty())
	fmt.Printf("Are motorcycle spaces full? %t\n", parkingLot.SpacesVehicleTypeFull(Motorcycle))
	fmt.Printf("Are car spaces full? %t\n", parkingLot.SpacesVehicleTypeFull(Car))
	fmt.Printf("Are van spaces full? %t\n", parkingLot.SpacesVehicleTypeFull(Van))
	fmt.Println("Available spaces for cars:", parkingLot.GetSpacesByVehicle(Car))
	fmt.Println("Available spaces for vans:", parkingLot.GetSpacesByVehicle(Van))
	fmt.Println("Available spaces for motorcycles:", parkingLot.GetSpacesByVehicle(Motorcycle))
	fmt.Println("Occupied spaces for cars:", parkingLot.GetOccupiedSpacesByVehicle(Car))
	fmt.Println("Occupied spaces for vans:", parkingLot.GetOccupiedSpacesByVehicle(Van))
	fmt.Println("Occupied spaces for motorcycles:", parkingLot.GetOccupiedSpacesByVehicle(Motorcycle))

}

// Método para criar um novo estacionamento
func NewParkingLot(totalSpaces int) *ParkingLot {
	prk := ParkingLot{
		TotalSpaces:         totalSpaces,
		OccupiedSpaces:      0,
		VehiclesType:        DistributeInitialSpaces(totalSpaces),
		VehiclesDescription: make(map[string]Vehicle),
	}

	return &prk
}

// Método para verificar se o estacionamento está cheio
func (p *ParkingLot) IsFull() bool {
	return p.OccupiedSpaces >= p.TotalSpaces
}

// Método para verificar se o estacionamento está vazio
func (p *ParkingLot) IsEmpty() bool {
	return p.OccupiedSpaces == 0
}

// Método para verificar se todas as vagas de um determinado tipo de veículo estão cheias
func (p *ParkingLot) SpacesVehicleTypeFull(vehicleType VehicleType) bool {
	return p.VehiclesType[vehicleType] == 0
}

// Método para obter o número de vagas restantes por veículo
func (p *ParkingLot) GetSpacesByVehicle(vehicleType VehicleType) int {
	return p.VehiclesType[vehicleType]
}

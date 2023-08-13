package internal

import "fmt"

// Método para distribuição inicial de vagas para cada tipo de veículo
func DistributeInitialSpaces(TotalSpaces int) map[VehicleType]int {
	vehiclesType := make(map[VehicleType]int)
	vehiclesType[Car] = TotalSpaces / 2
	vehiclesType[Motorcycle] = TotalSpaces / 4
	vehiclesType[Van] = TotalSpaces / 4

	return vehiclesType
}

// / Método para adicionar um novo veículo
func (p *ParkingLot) ParkVehicle(vehicle Vehicle) error {
	// Verifica se o estacionamento está cheio
	if p.OccupiedSpaces >= p.TotalSpaces {
		return fmt.Errorf("parking lot is full")
	}

	// Verifica se já há um veículo com a mesma placa no estacionamento
	for _, existsVehicle := range p.VehiclesDescription {
		if existsVehicle.LicensePlate == vehicle.LicensePlate {
			return fmt.Errorf("vehicle with the same license plate is already parked")
		}
	}

	if vehicle.Type == Van {
		// Verifica se há vagas de van disponíveis
		if p.VehiclesType[Van] > 0 {
			p.VehiclesType[Van]--
		} else if p.VehiclesType[Car] >= 3 {
			// Se não houver vaga de van disponível, estaciona em 3 vagas de carro
			p.VehiclesType[Car] -= 3
		} else {
			return fmt.Errorf("no available spaces for this vehicle type")
		}
	}

	if vehicle.Type == Car {
		// Verifica se há vagas de carro disponíveis
		if p.VehiclesType[Car] > 0 {
			p.VehiclesType[Car]--
		} else {
			return fmt.Errorf("no available spaces for cars")
		}
	}
	if vehicle.Type == Motorcycle {
		// Verifica se há vagas de moto disponíveis
		if p.VehiclesType[Motorcycle] > 0 {
			p.VehiclesType[Motorcycle]--
		} else {
			return fmt.Errorf("no available spaces for moto")
		}
	}

	// Registra o veículo no mapa de descrição de veículos
	p.VehiclesDescription[vehicle.LicensePlate] = vehicle
	p.OccupiedSpaces++

	// Retorna nil para indicar que o estacionamento ocorreu com sucesso
	return nil
}

// Método para remover um veículo do estacionamento
// func (p *ParkingLot) RemoveVehicle(licensePlate string) error {
// 	vehicle, exists := p.VehiclesDescription[licensePlate]
// 	if !exists {
// 		return fmt.Errorf("vehicle not found")
// 	}

// 	p.VehiclesType[vehicle.Type]++
// 	p.OccupiedSpaces--
// 	delete(p.VehiclesDescription, licensePlate)
// 	return nil
// }

// Método para obter o número de vagas ocupadas por veículo
func (p *ParkingLot) GetOccupiedSpacesByVehicle(vehicleType VehicleType) int {
	initialSpaces := p.TotalSpaces / 2
	if vehicleType == Motorcycle || vehicleType == Van {
		initialSpaces = p.TotalSpaces / 4
	}

	occupiedSpaces := initialSpaces - p.GetSpacesByVehicle(vehicleType)
	return occupiedSpaces
}

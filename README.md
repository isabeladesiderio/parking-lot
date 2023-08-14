
# Parking lot

O projeto é responsável por verificar ações dentro de um estacionamento.


## Configuração

- Necessário versão mínima 1.16 da linguagem Go.

## Pré-requisitos

Antes de usar a função `ParkingLotStatus`, é necessário preencher os seguintes elementos:

- O método `NewParkingLot` para criar um novo estacionamento, onde se insere o `totalSpaces` e retorna a estrutura `ParkingLot`.
- o método `ParkVehicle` para adicionar um novo veículo passando a estrutura `Vehicle`.

## Estrutura do Projeto

- Na pasta `cmd` está o arquivo `main.go` que possui a função `main()`, responsável pela execução do projeto.
- Na pasta `internal` contém o pacote `parkinglot`, onde está todo o código presente nesse projeto dividido em 4 arquivos:
    - `parking_test.go` é um arquivo onde contém cenários de testes voltados ao estacionamento.
    - `parking.go` é um arquivo onde se encontra todas as funções necessárias para a execução do projeto.
    - `vehicle_test.go` é um arquivo onde contém cenários de testes voltadas para os veículos.
    - `vehicle.go` é  um arquivo onde contém funções voltadas para os veículos.
## Testes

- Para rodar os testes execute o comando `go test -v ./...` no terminal. Esse comando serve para exibir informações detalhadas durante a execução dos testes em todos os pacotes e subpacotes do diretório atual.
- Para rodar a função main() basta apenas executar o comando `go run cmd/main.go`. Ela retornará o status do estacionamento.

## Descrevendo as funções

A função `ParkingLotStatus` segue os seguintes passos:

1.  A função `IsFull` tem como objetivo verificar se o estacionamento está cheio.
2.  A função `IsEmpty` tem como objetivo verificar se o estacionamento está vazio.
3.  A função `SpacesVehicleTypeFull` verifica se todas as vagas de um determinado tipo de veículo estão cheias.
4.  A função `GetSpacesByVehicle` tem como objetivo obter o número de vagas restantes por veículo.
5.  A função `GetOccupiedSpacesByVehicle` tem o objetivo de obter o número de vagas ocupadas por veículo.
6.  A função `DistributeSpaces` distribui o peso inicial de vagas para cada tipo de veículo.
7.  A função `ParkVehicle` retorna se as regras da verificação do estacionamento ocorreram com sucesso.


## Autor

- Nome: Isabela Desiderio
- E-mail: isabelaaraujodesiderio@gmail.com
- GitHub: isabeladesiderio
- LinkedIn: https://www.linkedin.com/in/isabela-araujo-desiderio/

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
- Na pasta `internal` cntém o pacote `parkinglot`, onde está todo o código presente nesse projeto dividido em 4 arquivos:
    - `parking_test.go` é um arquivo onde contém cenários de testes.
    - `parking.go` é um arquivo onde se encontra todas as funções necessárias para a execução do projeto.
    - `vehicle_test.go` é um arquivo onde contém funções voltadas para o cliente.
    - `vehicle.go` é  um arquivo onde contém funções voltadas para o Customer Success (CS).
## Testes

- Para rodar os testes execute o comando `go test -v ./...` no terminal. Esse comando serve para exibir informações detalhadas durante a execução dos testes em todos os pacotes e subpacotes do diretório atual.
- Para rodar a função main() basta apenas executar o comando `go run cmd/main.go`. Ela retornará o status do estacionamento.

## Descrevendo as funções

A função `ParkingLotStatus` segue os seguintes passos:

1.  A função `IsFull` tem como objetivo 


## Autor

- Nome: Isabela Desiderio
- E-mail: isabelaaraujodesiderio@gmail.com
- GitHub: isabeladesiderio
- LinkedIn: https://www.linkedin.com/in/isabela-araujo-desiderio/
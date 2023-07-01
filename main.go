package main

import (
	"banco/BankSimulationGolang/contas"
	"fmt"
)

func main() {
	contaDoPedro := contas.ContaCorrente{Titular: "Pedro", Saldo: 400}
	contaDaAna := contas.ContaCorrente{Titular: "Ana", Saldo: 600}

	status := contaDoPedro.Transferir(300, &contaDaAna)

	fmt.Println(status)

	fmt.Println(contaDoPedro)
	fmt.Println(contaDaAna)
}

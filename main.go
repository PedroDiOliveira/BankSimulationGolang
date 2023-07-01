package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, ValorDoBoleto float64) {
	conta.Sacar(ValorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(1000)
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis)
}

//utilizei do "go mod init nome da referencia" para conseguir conectar as pastas

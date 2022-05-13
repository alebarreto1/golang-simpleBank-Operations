package main

import (
	"banco1/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

//criar a interface varificarConta
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	// contaDoDenis := contas.ContaPoupanca{}
	// contaDoDenis.Depositar(100)
	// PagarBoleto(&contaDoDenis, 60)

	// fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())
}

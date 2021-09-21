package main

import (
	"fmt"
	"go_oo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)
	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(500)
	PagarBoleto(&contaDaLuiza, 400)
	fmt.Println(contaDaLuiza.ObterSaldo())
}

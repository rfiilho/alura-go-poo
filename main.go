package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Saque(valorBoleto)
}

type verificarConta interface {
	Saque(valor float64) string
}

func main() {
	contaRenato := contas.ContaCorrente{}
	contaRenato.Depositar(100)
	PagarBoleto(&contaRenato, 60)
	fmt.Println(contaRenato.GetSaldo())

	// clienteRenato := clientes.Titular{Nome: "Renato Filho", CPF: "123.123.123.12", Profissao: "Developer"}
	// contaRenato := contas.ContaCorrente{Titular: clienteRenato, Agencia: 123456, Conta: 123}
	// contaRenato.Depositar(1200)
	// fmt.Println(contaRenato.GetSaldo())

	// contaRenato := contas.ContaCorrente{Titular: "Renato", Saldo: 500}
	// contaMirela := contas.ContaCorrente{Titular: "Mirela", Saldo: 1000}

	// fmt.Println(contaRenato.Saldo)
	// fmt.Println(contaMirela.Saldo)

	// status := contaMirela.Transfer(200, &contaRenato)
	// if status {
	// 	fmt.Println(contaRenato.Saldo)
	// 	fmt.Println(contaMirela.Saldo)
	// }

}

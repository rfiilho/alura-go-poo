package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaRenato := contas.ContaCorrente{Titular: "Renato", Saldo: 500}
	contaMirela := contas.ContaCorrente{Titular: "Mirela", Saldo: 1000}

	fmt.Println(contaRenato.Saldo)
	fmt.Println(contaMirela.Saldo)

	status := contaMirela.Transfer(200, &contaRenato)
	if status {
		fmt.Println(contaRenato.Saldo)
		fmt.Println(contaMirela.Saldo)
	}

}

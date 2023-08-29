package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCorrente) Saque(valor float64) string {
	permitido := valor <= c.saldo && valor > 0
	if permitido {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Operação não realizada"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	permitido := valor > 0
	if permitido {
		c.saldo += valor
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Operação não realizada", c.saldo
	}
}

func (c *ContaCorrente) Transfer(valor float64, contaDestino *ContaCorrente) bool {
	permitido := valor <= c.saldo && valor > 0
	if permitido {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func main() {
	contaRenato := ContaCorrente{titular: "Renato", saldo: 500}
	contaMirela := ContaCorrente{titular: "Mirela", saldo: 1000}

	fmt.Println(contaRenato.saldo)
	fmt.Println(contaMirela.saldo)

	status := contaMirela.Transfer(200, &contaRenato)
	if status {
		fmt.Println(contaRenato.saldo)
		fmt.Println(contaMirela.saldo)
	}

	// status, valor := contaRenato.Depositar(500)
	// fmt.Println(status, valor)

}

package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64
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

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}

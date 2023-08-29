package contas

import (
	"banco/clientes"
)

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Saque(valor float64) string {
	permitido := valor <= c.saldo && valor > 0
	if permitido {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Operação não realizada"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	permitido := valor > 0
	if permitido {
		c.saldo += valor
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Operação não realizada", c.saldo
	}
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}

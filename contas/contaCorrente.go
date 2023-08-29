package contas

type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Saque(valor float64) string {
	permitido := valor <= c.Saldo && valor > 0
	if permitido {
		c.Saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Operação não realizada"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	permitido := valor > 0
	if permitido {
		c.Saldo += valor
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Operação não realizada", c.Saldo
	}
}

func (c *ContaCorrente) Transfer(valor float64, contaDestino *ContaCorrente) bool {
	permitido := valor <= c.Saldo && valor > 0
	if permitido {
		c.Saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

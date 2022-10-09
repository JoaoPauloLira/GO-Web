package entities

import (
	"fmt"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func (p Produto) ToString() string {
	return fmt.Sprintf("id: %d\nnome: %s\ndescricao: %spreco: %0.1f\nquantidade: %d", p.Id, p.Nome, p.Descricao, p.Preco, p.Quantidade)
}

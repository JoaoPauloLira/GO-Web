package business

import (
	"GOWeb/config"
	"GOWeb/entities"
	"GOWeb/repository"
	"database/sql"
	"fmt"
	"strconv"
)

func ObtemDb() *sql.DB {
	db, err := config.GetDB()

	if err != nil {
		panic(err.Error())
	}
	return db
}

func BuscarTodosProdutos() ([]entities.Produto, error) {

	produtos := []entities.Produto{}

	produtoModel := repository.ProdutoModel{Db: ObtemDb()}

	produtosAux, err2 := produtoModel.FindAll()
	if err2 != nil {
		panic(err2.Error())
	} else {
		for _, produto := range produtosAux {
			produtos = append(produtos, produto)
		}
	}

	return produtos, nil
}

func CriarNovoProduto(produto entities.Produto) {
	produtoModel := repository.ProdutoModel{Db: ObtemDb()}

	result, err2 := produtoModel.InserirNovotProduto(produto)

	if err2 != nil {
		panic(err2.Error())
	} else {
		println("Foram inseridos: ", result)
	}
}

func DeletarProduto(id string) {
	produtoModel := repository.ProdutoModel{Db: ObtemDb()}

	idProdutoInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	result, err2 := produtoModel.DeletarProduto(idProdutoInt)

	if err2 != nil {
		panic(err2.Error())
	} else {
		println("Foram excluidos: ", result, " Resgistros")
	}
}

func EditarProduto(id string, produto entities.Produto) {
	produtoModel := repository.ProdutoModel{Db: ObtemDb()}

	idProdutoInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	produto.Id = idProdutoInt

	fmt.Print("produto a ser editado", produto)

	result, err2 := produtoModel.EditarProduto(produto)

	if err2 != nil {
		panic(err2.Error())
	} else {
		println("Foram editados: ", result, " Resgistros")
	}
}

func BuscarProduto(id string) (entities.Produto) {
	produtoModel := repository.ProdutoModel{Db: ObtemDb()}

	idProdutoInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	result, err2 := produtoModel.BuscarProduto(idProdutoInt)

	if err2 != nil {
		panic(err2.Error())
	}

	return result
}

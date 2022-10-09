package repository

import (
	"GOWeb/entities"
	"database/sql"
	"fmt"
)

type ProdutoModel struct {
	Db *sql.DB
}

func (p ProdutoModel) FindAll() ([]entities.Produto, error) {
	rows, err := p.Db.Query("select * from produtos")

	var produtos []entities.Produto

	if err != nil {
		return nil, err
	} else {
		for rows.Next() {

			var id int
			var nome string
			var descricao string
			var preco float64
			var quantidade int

			err2 := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
			if err2 != nil {
				return nil, err2
			}

			produtoAux := entities.Produto{
				Id:         id,
				Nome:       nome,
				Descricao:  descricao,
				Preco:      preco,
				Quantidade: quantidade,
			}

			produtos = append(produtos, produtoAux)
		}
	}
	defer p.Db.Close()
	return produtos, nil
}

func (p ProdutoModel) InserirNovotProduto(produto entities.Produto) (int64, error) {

	//inserirDadosBanco, err := p.Db.Prepare(`insert into produtos (nome,descricao,preco,quantidade) values ($1, $2, $3, $4)`)
	//inserirDadosBanco.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	//
	//if err != nil {
	//	panic(err.Error())
	//}

	result, err := p.Db.Exec(`insert into produtos (nome,descricao,preco,quantidade) values (?, ?, ?, ?)`,
		produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)

	if err != nil {
		return 0, err
	} else {
		defer p.Db.Close()
		linhasAfetadas, _ := result.RowsAffected()
		return linhasAfetadas, nil
	}

}

func (p ProdutoModel) DeletarProduto(id int) (int64, error) {

	result, err := p.Db.Exec(`delete from produtos where id = ?`, id)

	if err != nil {
		return 0, err
	} else {

		linhasAfetadas, _ := result.RowsAffected()
		return linhasAfetadas, nil
	}
	defer p.Db.Close()
	return 0, nil
}

func (p ProdutoModel) EditarProduto(produto entities.Produto) (int64, error) {

	fmt.Print("produto a ser editado no repository", produto)

	result, err := p.Db.Exec(`update produtos 
									set nome = ?,
										descricao = ?,
										preco = ?,
										quantidade = ?
									where id = ?`,
		produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)

	if err != nil {
		return 0, err
	} else {

		linhasAfetadas, _ := result.RowsAffected()
		return linhasAfetadas, nil
	}
	defer p.Db.Close()
	return 0, nil
}

func (p ProdutoModel) BuscarProduto(id int) (entities.Produto, error) {
	rows, err := p.Db.Query("select * from produtos where id = ?", id)

	produto := entities.Produto{}

	if err != nil {
		panic(err.Error())
	} else {
		for rows.Next() {

			var id int
			var nome string
			var descricao string
			var preco float64
			var quantidade int

			err2 := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
			if err2 != nil {
				panic(err.Error())
			}

			produtoAux := entities.Produto{
				Id:         id,
				Nome:       nome,
				Descricao:  descricao,
				Preco:      preco,
				Quantidade: quantidade,
			}

			produto = produtoAux
		}
	}
	defer p.Db.Close()
	return produto, nil
}

package controllers

import (
	"GOWeb/business"
	"GOWeb/entities"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos, err := business.BuscarTodosProdutos()
	if err != nil {
		panic(err.Error())
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := entities.Produto{}
		produto = ConverterProdutoHtml(r)
		business.CriarNovoProduto(produto)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	business.DeletarProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cheguei ")
	idProduto := r.URL.Query().Get("id")
	fmt.Println("id = ", idProduto)
	produto := business.BuscarProduto(idProduto)

	fmt.Println("produto recuperado = ", produto)

	temp.ExecuteTemplate(w, "Editar", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idProduto := r.FormValue("id")

		produto := entities.Produto{}
		produto = ConverterProdutoHtml(r)
		business.EditarProduto(idProduto, produto)
	}
	http.Redirect(w, r, "/", 301)
}

func ConverterProdutoHtml(r *http.Request) (p entities.Produto) {
	p.Nome = r.FormValue("nome")
	p.Descricao = r.FormValue("descricao")
	precoAux := r.FormValue("preco")
	quantidadeAux := r.FormValue("quantidade")

	precoConvertido, err := strconv.ParseFloat(precoAux, 64)
	if err != nil {
		log.Println("Erro ao tentar converter o preço: ", err.Error())
	}
	p.Preco = precoConvertido

	quantidadeConvertida, err := strconv.Atoi(quantidadeAux)
	if err != nil {
		log.Println("Erro ao tentar converter a quantidade: ", err.Error())
	}
	p.Quantidade = quantidadeConvertida
	return
}

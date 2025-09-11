package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Estrutura de Usuario
type Usuario struct {
	Id   int    `json:"id"`   // mapeia o atributo id para o json
	Nome string `json:"nome"` // mapeia o atributo nome para o json

}

// UsuarioHandler analisa o request e delega para a função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	stringid := strings.TrimPrefix(r.URL.Path, "/usuarios/") // remove o prefixo /usuarios/ da url considerando que a url é /usuarios/id
	id, _ := strconv.Atoi(stringid)                          // converte a string id para um inteiro

	switch {
	case r.Method == "GET" && stringid >= "0":
		usuarioPorId(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	// case r.Method == "POST":
	// 	usuarioInserir(w, r)
	// case r.Method == "PUT":
	// 	usuarioAtualizar(w, r, id)
	// case r.Method == "DELETE":
	// 	usuarioDeletar(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Método não permitido") // gera um erro 404
	}
}
func usuarioPorId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:338080db@tcp(localhost:3306)/cursogo")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao abrir conexão com o banco de dados")
	}
	defer db.Close() // fecha a conexão com o banco de dados ao final da função

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.Id, &u.Nome) //Query Row retorna uma única linha da consulta

	json, _ := json.Marshal(u) // converte o usuário para um json

	w.Header().Set("Content-Type", "application/json") //o tipo de dado que será retornado é json
	fmt.Fprint(w, string(json))                        // escreve o json no response writer
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:338080db@tcp(localhost:3306)/cursogo")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao abrir conexão com o banco de dados")
	}
	defer db.Close() // fecha a conexão com o banco de dados ao final da função

	rows, _ := db.Query("select id, nome from usuarios") //retorna todas as linhas da consulta

	var usuarios []Usuario // cria um slice de usuários
	for rows.Next() {      //pega cada elemento da linha e atribui ao usuário
		var u Usuario
		rows.Scan(&u.Id, &u.Nome)      // atribui o id e o nome do usuário ao usuário
		usuarios = append(usuarios, u) // adiciona o usuário ao slice de usuários
	}

	json, _ := json.Marshal(usuarios)                  // converte o slice de usuários para um json
	w.Header().Set("Content-Type", "application/json") //seta o Header do response writer que será retornado para tipo json
	fmt.Fprintf(w, string(json))                       // escreve o json no response writer
}

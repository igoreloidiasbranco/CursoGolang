package main

//import go get github.com/go-sql-driver/mysql
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	//abre conexão com o banco de dados
	fmt.Println("Conectando ao MySQL...")
	db, err := sql.Open("mysql", "root:338080db@tcp(localhost:3306)/")
	if err != nil {
		fmt.Printf("Erro ao abrir conexão: %v\n", err)
		panic(err)
	}
	defer db.Close() // fecha a conexão com o banco de dados ao final do método main

	// Testa a conexão
	fmt.Println("Testando conexão...")
	err = db.Ping()
	if err != nil {
		fmt.Printf("Erro ao conectar com o banco: %v\n", err)
		panic(err)
	}
	fmt.Println("Conexão estabelecida com sucesso!")

	//cria bd

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, `create table if not exists usuarios (
		id integer auto_increment,
		nome varchar(80),
		primary key (id)
	)`)
}

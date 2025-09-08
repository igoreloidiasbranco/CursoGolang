package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Conectando ao MySQL...")
	db, err := sql.Open("mysql", "root:338080db@tcp(localhost:3306)/cursogo")
	if err != nil {
		fmt.Printf("Erro ao abrir conexão: %v\n", err)
		panic(err)
	}
	defer db.Close()

	// Testa a conexão
	fmt.Println("Testando conexão...")
	err = db.Ping()
	if err != nil {
		fmt.Printf("Erro ao conectar com o banco: %v\n", err)
		panic(err)
	}
	fmt.Println("Conexão estabelecida com sucesso!")

	stmt, err := db.Prepare("insert into usuarios (nome) values (?)") // prepara o statement para inserir os usuários
	if err != nil {
		fmt.Printf("Erro ao preparar statement: %v\n", err)
		panic(err)
	}
	defer stmt.Close() // fecha o statement ao final do método main

	fmt.Println("Inserindo usuários...")

	_, err = stmt.Exec("Maria")
	if err != nil {
		fmt.Printf("Erro ao inserir Maria: %v\n", err)
	} else {
		fmt.Println("Maria inserida com sucesso!")
	}

	_, err = stmt.Exec("João")
	if err != nil {
		fmt.Printf("Erro ao inserir João: %v\n", err)
	} else {
		fmt.Println("João inserido com sucesso!")
	}

	res, err := stmt.Exec("Pedro")
	if err != nil {
		fmt.Printf("Erro ao inserir Pedro: %v\n", err)
	} else {
		id, _ := res.LastInsertId()
		fmt.Printf("Pedro inserido com sucesso! ID: %d\n", id)
	}
}

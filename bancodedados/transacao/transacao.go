package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // importa o driver de conexão com o banco de dados "mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:338080db@tcp(localhost:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // fecha a conexão com o banco de dados ao final do método main

	//Transação, como se fosse o Transaction do Java, ou tudo acontece ou nada acontece.
	tx, _ := db.Begin()                                                    //inicia e cria um objeto tx(transação) para fornecer os comandos de sql
	stmt, _ := tx.Prepare("insert into usuarios (id, nome) values (?, ?)") // prepara o statement para inserir os usuários

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // vai gerar um erro pois o id 1 já existe

	if err != nil {
		tx.Rollback() // se der erro, desfaz a transação
		log.Fatal(err)
	}

	tx.Commit() // se não der erro, commita a transação
}

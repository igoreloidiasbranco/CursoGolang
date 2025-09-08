package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:338080db@tcp(localhost:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // fecha a conexão com o banco de dados ao final do método main

	//Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?") // prepara o statement para atualizar o nome do usuário

	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Valeska", 2)

	//Delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}

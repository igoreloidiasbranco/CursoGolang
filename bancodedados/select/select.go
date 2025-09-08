package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:338080db@tcp(localhost:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // fecha a conexão com o banco de dados ao final do método main

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 5) // seleciona todos os usuários com id maior que 5
	defer rows.Close()                                                   // fecha as linhas ao final do método main

	for rows.Next() { //pega cada elemento da linha e atribui ao usuário
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}

package main

import "database/sql"

type Estrutura struct {
	v1 int
	v2 string
}

func main() {
	db, err := sql.Open("sqlite3", "t1.db")
	if err != nil {
		panic(err)
	}

	// Ultima coisa a ser executado
	defer db.Close()

	test := Estrutura{v1: 1, v2: "test"}
	insertValue(db, test)
}

func insertValue(db *sql.DB, value Estrutura) {
	_, err := db.Exec("INSERT INTO values (v1, v2) VALUES (?, ?)", value.v1, value.v2)
	if err != nil {
		panic(err)
	}
}

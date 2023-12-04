package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

type Repository interface {
	GetbyID(id int)
}

func GetDB(sc string) *DB {
	db, err := sql.Open("sqlite3", sc)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db}
}

func InsertMessage() {

}

func (d DB) GetByID(id int) []Message {
	mes := []Message{}
	rows, err := d.Query("SELECT * FROM notify WHERE id=?;", id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		m := Message{}
		err = rows.Scan(&m.Id, &m.Message, &m.Date, &m.Time)
		if err != nil {
			log.Fatal(err)
		}
		mes = append(mes, m)
	}
	return mes
}

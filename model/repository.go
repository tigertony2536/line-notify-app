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

func (db *DB) InsertNotification(noti Notification) (int, error) {
	result, err := db.Exec(`INSERT INTO notify(message, date, time) VALUES(?,?,?);`, noti.Message, noti.Date, noti.Time)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (db *DB) GetByID(id int) (Notification, error) {
	row := db.QueryRow("SELECT * FROM notify WHERE id=?;", id)

	m := Notification{}
	err := row.Scan(&m.ID, &m.Message, &m.Date, &m.Time)
	if err != nil {
		return Notification{}, err
	}
	return m, nil
}

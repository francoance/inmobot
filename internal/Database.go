package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const databaseName string = "inmobot.db"

var db *sql.DB

func InitDatabase() {
	db, _ = sql.Open("sqlite3", databaseName)

	db.Exec(`CREATE TABLE IF NOT EXISTS Inmueble (
	    kid VARCHAR(255),
		superficie_total VARCHAR(255), 
		superficie_cubierta VARCHAR(255),
		precio VARCHAR(255),
		foto VARCHAR(255),
		habitaciones VARCHAR(255),
		banos VARCHAR(255),
    	url VARCHAR(255),
    	direccion VARCHAR(255)
	);`)
}

func InmuebleExists(kid string) bool {
	rows, _ := db.Query(fmt.Sprintf("SELECT count(kid) as count FROM Inmueble where kid = %v", kid))
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	return count > 0
}

func SaveInmueble(inm Inmueble) {
	statement, _ := db.Prepare("INSERT INTO Inmueble (kid, superficie_total, superficie_cubierta, precio, foto, habitaciones, banos, url, direccion) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	statement.Exec(inm.kid, inm.supTotal, inm.supCubierta, inm.precio, inm.foto, inm.habitaciones, inm.banos, inm.url, inm.direccion)
}

func DatabaseExists() bool {
	return true
}

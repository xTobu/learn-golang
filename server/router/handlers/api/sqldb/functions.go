package sqldb

import (
	"database/sql"

	"log"
)

func dbGetConn() *sql.DB {
	db, err := sql.Open("mssql", ConnectionStr)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	return db
}

// DBGetStudents 123
func DBGetStudents() (data []Student, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("select * from [vueStudent]")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var student Student
		rows.Scan(&student.ID, &student.Name, &student.Email, &student.CreatedTime)
		data = append(data, student)
		count++
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	return data, count
}

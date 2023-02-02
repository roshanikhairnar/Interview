package application

import (
	"database/sql"
	"fmt"
)

type Student struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"LastName"`
	Department string `json:"department"`
}

func (student *Student) getstudent(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * from studentData WHERE id=%d", student.Id)
	return db.QueryRow(statement).Scan(&student.Id, &student.FirstName, &student.LastName, &student.Department)
}
func getstudents(db *sql.DB) ([]Student, error) {
	statement := fmt.Sprintf("SELECT * from studentData")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	students := []Student{}

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Department); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (student *Student) updatestudent(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE studentData SET firstName='%s'WHERE id=%d", student.FirstName, student.Id)
	_, err := db.Exec(statement)
	return err
}

func (student *Student) deletestudent(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM studentData WHERE id=%d", student.Id)
	_, err := db.Exec(statement)
	return err
}

func (student *Student) createstudent(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO studentData(id,firstName,lastName,department) VALUES('%d','%s','%s','%s')", student.Id, student.FirstName, student.LastName, student.Department)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_studentID()").Scan(&student.studentID)

	if err != nil {
		return err
	}

	return nil
}

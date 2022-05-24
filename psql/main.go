package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

const (
	DB_HOST = "localhost"
	DB_PORT = 5432
	DB_USER = "ktiarad"
	DB_PASS = "123456"
	DB_NAME = "learn_postgres"
)

var db *sql.DB

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	// CREATE EMPLOYEE
	// emp := Employee{
	// 	Email:    "naruto@konoha.com",
	// 	FullName: "Uzumaki Naruto",
	// 	Age:      21,
	// 	Division: "Developer",
	// }

	// err = createEmployee(db, &emp)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// }

	// UPDATE EMPLOYEE
	// err = updateEmployee(db)
	// if err != nil {
	// 	fmt.Println("error :", err)
	// 	return
	// }

	// GET EMPLOYEE BY ID
	// employee, err := getEmployeeById(db, 2)
	// if err != nil {
	// 	fmt.Println("error:", err.Error())
	// }
	// employee.Print()

	// DELETE EMPLOYEE BY ID
	err = deleteEmployeeById(db, 1)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	// GET ALL EMPLOYEES
	employees, err := getallEmployees(db)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	for _, employee := range *employees {
		employee.Print()
	}

}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}

func getallEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var employees []Employee

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return &employees, nil
}

func (e *Employee) Print() {
	fmt.Println("ID :", e.ID)
	fmt.Println("FullName :", e.FullName)
	fmt.Println("Email :", e.Email)
	fmt.Println("Age :", e.Age)
	fmt.Println("Division :", e.Division)
	fmt.Println("==========")

}

func createEmployee(db *sql.DB, request *Employee) error {
	query := `
		INSERT INTO employees(full_name, email, age, division)
		VALUES($1, $2, $3, $4)
	`
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()
}

func getEmployeeById(db *sql.DB, id int) (*Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees 
		WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var emp Employee

	err = row.Scan(
		&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division,
	)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func updateEmployee(db *sql.DB) error {
	query := `
		UPDATE employees
		SET full_name = $2, email = $3, division = $4, age = $5
		WHERE id = $1
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := db.Exec(query, 1, "Uchiha Sasuke", "sasuke@konoha.com", "Product Manager", 30)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Total updated data:", count)

	return nil

}

func deleteEmployeeById(db *sql.DB, id int) error {
	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Data with id %d has been deleted, %d rows affected\n", id, count)
	return nil
}

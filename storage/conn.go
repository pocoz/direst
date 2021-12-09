package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/dieqnt/direst/models"
)

type Engine struct {
	db *sql.DB
}

func New() (*Engine, error) {
	engine, err := connect()
	if err != nil {
		return nil, err
	}

	return engine, nil
}

func connect() (*Engine, error) {
	const (
		host     = "localhost"
		port     = 5432
		userName = "pocoz"
		password = ""
		dbname   = "pocoz"
	)

	// connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		userName,
		password,
		dbname,
	)

	// open database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("psql connected successfully")

	engine := &Engine{
		db: db,
	}

	return engine, err
}

func (e *Engine) UserInsert(user *models.User) error {
	query := `insert into "users"("email", "password") values($1, $2)`
	_, err := e.db.Exec(query, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) UserUpdate(user *models.User) error {
	query := `update "users" set "email"=$1, "password"=$2 where "id"=$3`
	_, err := e.db.Exec(query, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) UserDelete(id int) error {
	fmt.Println(id)

	query := `delete from "users" where id=$1`
	_, err := e.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) UserGetList() ([]*models.User, error) {
	rows, err := e.db.Query(`SELECT "id", "email", "password" FROM "users"`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userList := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}

		userList = append(userList, user)
	}

	return userList, nil
}

func (e *Engine) UserGetByID(id int) (*models.User, error) {
	row := e.db.QueryRow(`SELECT "id", "email", "password" FROM "users" where "id"=$1`, id)

	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

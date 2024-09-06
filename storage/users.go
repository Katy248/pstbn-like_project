package storage

import (
	m "cdecode/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func getUserFromRow(row *sql.Rows) (*m.User, error) {
	user := &m.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Password, &user.IsAdmin); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers(db *sql.DB) []*m.User {
	q := `SELECT * FROM users`
	row, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	users := []*m.User{}
	for row.Next() {
		user, err := getUserFromRow(row)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	return users
}

func GetUserById(db *sql.DB, id int) (*m.User, error) {
	query := "SELECT * FROM users WHERE id_user = ?"
	stmt, err := db.Prepare(query)

	if err != nil {
		return nil, err
	}

	row, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		user, err := getUserFromRow(row)
		return user, err
	}
	return nil, errors.New(fmt.Sprintf("No such user with id %d", id))
}

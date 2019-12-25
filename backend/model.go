package main

import (
	"database/sql"
)

type user struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (p *user) getUser(db *sql.DB) error {
	return db.QueryRow("SELECT firstname, lastname FROM users WHERE id=$1",
		p.ID).Scan(&p.FirstName, &p.LastName)
}

func (p *user) updateUser(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE users SET firstname=$1, lastname=$2 WHERE id=$3",
			p.FirstName, p.LastName, p.ID)

	return err
}

func (p *user) deleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", p.ID)

	return err
}

func (p *user) createUser(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO users(firstname, lastname) VALUES($1, $2) RETURNING id",
		p.FirstName, p.LastName).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	rows, err := db.Query(
		"SELECT id, firstname, lastname FROM users LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var p user
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	return users, nil
}

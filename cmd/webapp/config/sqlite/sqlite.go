package sqlite

import (
	"database/sql"
	"errors"
	"os"

	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tranthienphucbreak/parsleytest/internal/patient"
)

type sqliteService struct {
	conn *sql.DB
}

func NewSqliteService() *sqliteService {
	var db *sql.DB
	var err error
	_, err = os.Open("./database.db")
	if err != nil {
		db, err = sql.Open("sqlite3", "./config/sqlite/database.db")
	} else {
		db, err = sql.Open("sqlite3", "./database.db")
	}
	if err != nil {
		panic(err)
	}
	return &sqliteService{db}
}

func (s *sqliteService) Exec(query string, params []interface{}) (int64, error) {
	rs, err := s.conn.Exec(query, params...)
	if err != nil {
		return 0, err
	}
	rowAffected, err := rs.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowAffected == 0 {
		return 0, errors.New(patient.ERROR_ITEM_NOT_FOUND)
	}
	return rowAffected, err
}

func (s *sqliteService) Query(query string, params []interface{}) ([]patient.Person, error) {
	rs := []patient.Person{}
	rows, err := s.conn.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		person := patient.Person{}
		rows.Scan(&person.ID, &person.FirstName, &person.MiddleName, &person.LastName, &person.Email, &person.DOB, &person.Gender, &person.Status,
			&person.TermsAccepted, &person.TermsAcceptedAt, &person.AddressStreet, &person.AddressCity, &person.AddressState, &person.AddressZip, &person.Phone)
		rs = append(rs, person)
	}
	return rs, nil
}

func (s *sqliteService) QueryRow(query string, params []interface{}) (*patient.Person, error) {
	row := s.conn.QueryRow(query, params...)
	person := patient.Person{}
	row.Scan(&person.ID, &person.FirstName, &person.MiddleName, &person.LastName, &person.Email, &person.DOB, &person.Gender, &person.Status,
		&person.TermsAccepted, &person.TermsAcceptedAt, &person.AddressStreet, &person.AddressCity, &person.AddressState, &person.AddressZip, &person.Phone)
	if person.ID == "" {
		return nil, errors.New(patient.ERROR_ITEM_NOT_FOUND)
	}
	return &person, nil
}

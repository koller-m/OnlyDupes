package models

import (
	"database/sql"
	"time"
)

type Dupe struct {
	ID      int
	Dupe    string
	Content string
	Created time.Time
	Expires time.Time
}

type DupeModel struct {
	DB *sql.DB
}

func (m *DupeModel) Insert(dupe string, content string, expires int) (int, error) {
	stmt := `INSERT INTO dupes (dupe, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, dupe, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *DupeModel) Get(id int) (*Dupe, error) {
	return nil, nil
}

func (m *DupeModel) Latest() ([]*Dupe, error) {
	return nil, nil
}

package models

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, dupe, content, created, expires FROM dupes
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &Dupe{}

	err := row.Scan(&s.ID, &s.Dupe, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *DupeModel) Latest() ([]*Dupe, error) {
	stmt := `SELECT id, dupe, content, created, expires FROM dupes
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dupes := []*Dupe{}

	for rows.Next() {
		s := &Dupe{}

		err = rows.Scan(&s.ID, &s.Dupe, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		dupes = append(dupes, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dupes, nil
}

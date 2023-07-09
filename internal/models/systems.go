package models

import (
	"database/sql"
	"errors"
	"time"
)

type System struct {
  ID int
  Name string
  Content string
  IsArchived bool
  CreatedAt time.Time
}

type SystemModel struct {
  DB *sql.DB
}

func (m *SystemModel) Insert(name string, content string) (int, error) {
  stmt := `INSERT INTO systems (name, content)
  VALUES (?, ?)`

  result, err := m.DB.Exec(stmt, name, content)
  if err != nil {
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }

  return int(id), err
}

func (m *SystemModel) Get(id int) (*System, error) {
  s := &System{}

  stmt := `SELECT id, name, content, is_archived, created_at FROM systems
  WHERE is_archived = false AND id = ?`

  err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Name, &s.Content, &s.IsArchived, &s.CreatedAt)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, ErrNoRecord
    } else {
      return nil, err
    }
  }

  return s, nil
}

func (m *SystemModel) List() ([]*System, error) {
  stmt := `SELECT id, name FROM systems
  WHERE is_archived = false ORDER BY id DESC`

  rows, err := m.DB.Query(stmt)
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  systems := []*System{}

  for rows.Next() {
    s := &System{}

    err = rows.Scan(&s.ID, &s.Name)
    if err != nil {
      return nil, err
    }

    systems = append(systems, s)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }

  return systems, nil
}

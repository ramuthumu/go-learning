package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

type Ticket struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
type Store struct{ db *sql.DB }

func openStore(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	// One connection keeps this small SQLite lesson simple. Revisit pooling later.
	db.SetMaxOpenConns(1)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tickets (
        id INTEGER PRIMARY KEY, title TEXT NOT NULL, done INTEGER NOT NULL DEFAULT 0
    )`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("create schema: %w", err)
	}
	return &Store{db: db}, nil
}

func (s *Store) Create(ctx context.Context, title string) (Ticket, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO tickets(title) VALUES (?)", title)
	if err != nil {
		return Ticket{}, fmt.Errorf("create ticket: %w", err)
	}
	id, err := result.LastInsertId()
	return Ticket{ID: id, Title: title}, err
}

func (s *Store) List(ctx context.Context) ([]Ticket, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, title, done FROM tickets ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("list tickets: %w", err)
	}
	defer rows.Close()
	tickets := make([]Ticket, 0)
	for rows.Next() {
		var t Ticket
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, rows.Err()
}

func (s *Store) Complete(ctx context.Context, id int64) error {
	result, err := s.db.ExecContext(ctx, "UPDATE tickets SET done=1 WHERE id=?", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

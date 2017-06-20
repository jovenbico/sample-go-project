package model

import "database/sql"

// Model model type
type Model interface {
	Save(tx *sql.Tx) error
	Update(tx *sql.Tx) error
	Delete(tx *sql.Tx) error
}

package bnrsqlx

import "time"

type IDModel struct {
	ID string `db:"id"`
	BaseModel
}

type IDModelNoDeleted struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type IDIntModel struct {
	ID int64 `db:"id"`
	BaseModel
}

type BaseModel struct {
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

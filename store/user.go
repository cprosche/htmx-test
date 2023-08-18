package store

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id" form:"id"`
	Email     string    `db:"email" json:"email" form:"email"`
	Password  string    `db:"password" json:"password" form:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at" form:"created_at`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" form:"updated_at`
}

// TODO: add password hashing
func (u *User) Create(db *sqlx.DB) error {
	id := uuid.New()
	time := time.Now()

	_, err := db.Exec(`
	INSERT INTO 
		users (id, email, password, created_at, updated_at) 
	VALUES 
		($1, $2, $3, $4, $5)
	`,
		id,
		u.Email,
		u.Password,
		time,
		time,
	)
	if err != nil {
		return err
	}

	u.ID = id
	u.CreatedAt = time
	u.UpdatedAt = time

	return nil
}

func (u *User) Login(db *sqlx.DB) error {
	err := db.Get(u, `
	SELECT * FROM users WHERE email = $1 AND password = $2
	`,
		u.Email,
		u.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

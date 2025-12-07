package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/paralaxrus/health-project/dbsvc/internal/storage/model"
)

var ErrUserNotFound = errors.New("user not found")

type UserDataSource struct {
	db IDatabase
}

func NewUserDataSource() *UserDataSource {
	return &UserDataSource{}
}

type Index struct {
	id    int64
	email string
}

func NewIndex(id int64, email string) Index {
	return Index{id: id, email: email}
}

func (s *UserDataSource) Connect() error {
	url := os.Getenv("DATABASE_URL")
	log.Printf("connecting to database: %s", url)
	if len(url) == 0 {
		return fmt.Errorf("health database url is not set")
	}
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	s.db, err = NewPostgresDatabase(pool)
	if err != nil {
		return fmt.Errorf("Unable to create postgres db %v", err)
	}
	err = s.createSchema()
	success := "completed"
	if err != nil {
		success = "failed"
	}
	log.Printf("connection %s", success)
	return err
}

func (s *UserDataSource) Disconnect() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *UserDataSource) CreateUser(ctx context.Context, name string, email string, password string) (int, error) {
	const query = `
    	INSERT INTO users (email, name, created_at, password)
    	VALUES ($1, $2, $3, $4)
    	RETURNING id
	`
	user := model.NewUser(name, email, password)
	var id int
	err := s.db.QueryRow(ctx, query, user.Email, user.Name, user.Created, user.Password).Scan(&id)
	return id, err
}

func (s *UserDataSource) FindUser(ctx context.Context, index Index) (*model.User, error) {
	const query = "SELECT * FROM users WHERE email = $1 LIMIT 1"

	row := s.db.QueryRow(ctx, query, index.email)

	var id int
	var user model.User
	err := row.Scan(&id, &user.Email, &user.Name, &user.Created, &user.Password)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("user %s %w", index.email, ErrUserNotFound)
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserDataSource) createSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		password TEXT NOT NULL
	);
	`

	log.Printf("executing create users table query")
	log.Printf("connect completed, db=%v", s.db)
	_, err := s.db.Exec(context.Background(), query)
	log.Printf("done executing create users table query, err=%v", err)
	return err
}

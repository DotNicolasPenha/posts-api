package user

import (
	"context"
	"time"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *Repository {
	if conn == nil {
		logger.Fatal("connection of repository user is nil")
	}
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) Insert(createUserDto CreateUserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.conn.Exec(
		ctx, "INSERT INTO users (nameuser,passwordhash,bio) VALUES ($1,$2,$3)",
		createUserDto.Username, createUserDto.Password, createUserDto.Bio,
	)
	return err
}
func (r *Repository) FindMany() ([]UserPreviewPerson, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := r.conn.Query(ctx, "SELECT nameuser,bio,created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Users []UserPreviewPerson
	for rows.Next() {
		var User UserPreviewPerson
		if err := rows.Scan(&User.Username, &User.Bio, &User.Created_At); err != nil {
			return nil, err
		}
		Users = append(Users, User)
	}
	return Users, nil
}
func (r *Repository) FindByNameUser(nameuser string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err := r.conn.QueryRow(ctx, "SELECT nameuser,passwordhash FROM users WHERE nameuser=$1", nameuser).Scan(
		&user.Username, &user.PasswordHash,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

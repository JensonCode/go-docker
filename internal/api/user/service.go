package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JensonCode/go-docker/internal/models"
	"github.com/JensonCode/go-docker/pkg/bcrypt"
	"github.com/JensonCode/go-docker/pkg/database"
)

type UserServices struct{}

var UserService = new(UserServices)

func scanUsers(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServices) Create(req *models.UserRequest) (*models.User, error) {

	used, err := s.IsUsernameExist(req.Username)
	if err != nil {
		return nil, err
	}
	if used && req.Username != "admin" {
		return nil, fmt.Errorf("username has been used")
	}

	encrypted, err := bcrypt.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = encrypted

	ok := s.Save(req)

	if !ok {
		return nil, fmt.Errorf("error saving user")
	}

	return &models.User{
		Username:  req.Username,
		Password:  encrypted,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (s *UserServices) GetByID(id uint) ([]*models.User, error) {

	query := `select * from users where id=$1`

	rows, err := database.Postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*models.User{}
	for rows.Next() {
		user, err := scanUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserServices) IsUsernameExist(username string) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE username = $1`

	err := database.Postgres.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking username existence: %w", err)
	}

	return count > 0, nil
}

func (s *UserServices) Save(req *models.UserRequest) (ok bool) {
	query := `insert into users
	(username, password, created_at)
	values ($1, $2, $3)`

	_, err := database.Postgres.DB.Query(query,
		req.Username,
		req.Password,
		time.Now(),
	)

	return err == nil
}

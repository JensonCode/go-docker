package user

import (
	"fmt"
	"time"

	"github.com/JensonCode/go-docker/internal/models"
	"github.com/JensonCode/go-docker/pkg/bcrypt"
	"github.com/JensonCode/go-docker/pkg/database"
)

type UserServices struct{}

var UserService = new(UserServices)

func (s *UserServices) Create(req *models.CreateUserRequest) (*models.User, error) {

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

func (s *UserServices) Update(req *models.UpdateUserRequest) (*models.User, error) {

	user, err := s.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	compare := bcrypt.CompareHashAndPassword(user.Password, req.Password)
	if !compare {
		return nil, fmt.Errorf("incorrect password")
	}

	encrypted, err := bcrypt.EncryptPassword(req.NewPassword)
	if err != nil {
		return nil, err
	}

	req.Password = encrypted

	ok := s.Save(&req.CreateUserRequest)

	if !ok {
		return nil, fmt.Errorf("error saving user")
	}

	return &models.User{
		Username:  req.Username,
		Password:  encrypted,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (s *UserServices) GetByUsername(username string) (*models.User, error) {

	user := new(models.User)

	query := `select * from users where username=$1 limit 1`

	err := database.Postgres.DB.QueryRow(query, username).Scan(
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

func (s *UserServices) IsUsernameExist(username string) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE username = $1`

	err := database.Postgres.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking username existence: %w", err)
	}

	return count > 0, nil
}

func (s *UserServices) Save(req *models.CreateUserRequest) (ok bool) {
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

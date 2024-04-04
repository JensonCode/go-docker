package user

import (
	"database/sql"
	"time"

	"github.com/JensonCode/go-docker/internal/models"
	"github.com/JensonCode/go-docker/pkg/bcrypt"
	"github.com/JensonCode/go-docker/pkg/database"
	"github.com/JensonCode/go-docker/pkg/request"
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

func (s *UserServices) CreateUser(req *request.UserRequest) (*models.User, error) {

	encrypted, err := bcrypt.EncryptPassword(req.Password)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Username:  req.Username,
		Password:  encrypted,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (s *UserServices) GetUser() ([]*models.User, error) {

	query := `select * from users`

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

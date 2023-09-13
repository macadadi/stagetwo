package repository

import (
	"context"
	"errors"
	"fmt"
	"macadadi/stagetwo/db"
	"macadadi/stagetwo/form"
	"macadadi/stagetwo/model"
)

const (
	Adduser     = "INSERT INTO users (name) VALUES ($1)"
	getAllUsers = "SELECT id, name  FROM users"
	// Add the following
	deleteuser = "DELETE FROM users WHERE id= $1"
	findById   = getAllUsers + " WHERE id= $1"
	updateUser = "UPDATE users SET name = $1 WHERE id =$2"
)

type (
	UserInterface interface {
		AddUser(ctx context.Context, db db.DB, form *form.User) error
		GetAllUsers(ctx context.Context, db db.DB) ([]*model.User, error)
		UpdateUser(ctx context.Context, db db.DB, form *form.User) error
		DeleteUser(ctx context.Context, db db.DB, id int64) error
		FindUserByID(ctx context.Context, db db.DB, id int64) (*model.User, error)
	}
	UserRepository struct{}
)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) AddUser(ctx context.Context, db db.DB, form *form.User) error {
	_, err := db.ExecContext(ctx, Adduser, &form.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserRepository) GetAllUsers(ctx context.Context, db db.DB) ([]*model.User, error) {
	rows, err := db.QueryContext(ctx, getAllUsers)

	if err != nil {
		return []*model.User{}, err
	}
	var users = make([]*model.User, 0)
	for rows.Next() {
		var user model.User

		rows.Scan(
			&user.Id,
			&user.Name,
		)
		if err != nil {
			return []*model.User{}, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *UserRepository) FindUserByID(ctx context.Context, db db.DB, id int64) (*model.User, error) {
	var user model.User
	row := db.QueryRowContext(ctx, findById, id)
	err := row.Scan(
		&user.Id,
		&user.Name,
	)
	if err != nil {
		return &model.User{}, errors.New(err.Error())
	}
	return &user, nil
}

func (s *UserRepository) DeleteUser(ctx context.Context, db db.DB, id int64) error {

	_, err := s.FindUserByID(ctx, db, id)

	if err != nil {
		fmt.Print(err.Error(), "we got this error", id, "cur id")
		return errors.New("User could not be found")
	}
	_, err = db.ExecContext(ctx, deleteuser, id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (s *UserRepository) UpdateUser(ctx context.Context, db db.DB, form *form.User) (*model.User, error) {
	_, err := db.ExecContext(ctx, updateUser, form.Name, form.Id)
	if err != nil {
		return &model.User{}, errors.New("could not update user")
	}
	user, err := s.FindUserByID(ctx, db, form.Id)

	if err != nil {
		return &model.User{}, errors.New("could not find user")
	}
	return user, nil

}

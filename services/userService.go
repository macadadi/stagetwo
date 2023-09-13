package services

import (
	"context"
	"errors"
	"fmt"
	"macadadi/stagetwo/db"
	"macadadi/stagetwo/form"
	"macadadi/stagetwo/model"
	"macadadi/stagetwo/repository"
)

type (
	UserServiceInterface interface {
		AddUser(ctx context.Context, db db.DB, form *form.User) error
		GetAllUsers(ctx context.Context, db db.DB) ([]*model.User, error)
		UpdateUser(ctx context.Context, db db.DB, form *form.User) error
		DeleteUser(ctx context.Context, db db.DB, id int64) error
		FindUserByID(ctx context.Context, db db.DB, id int64) (*model.User, error)
	}

	UserService struct {
		UserRepository *repository.UserRepository
	}
)

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (s *UserService) AddUser(ctx context.Context, db db.DB, form *form.User) error {
	err := s.UserRepository.AddUser(ctx, db, form)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context, db db.DB) ([]*model.User, error) {
	users, err := s.UserRepository.GetAllUsers(ctx, db)

	if err != nil {
		return []*model.User{}, err
	}
	return users, nil
}
func (s *UserService) FindUserByID(ctx context.Context, db db.DB, id int64) (*model.User, error) {
	user, err := s.UserRepository.FindUserByID(ctx, db, id)

	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, db db.DB, id int64) error {

	err := s.UserRepository.DeleteUser(ctx, db, id)
	if err != nil {
		return errors.New(fmt.Sprintf("current error %v ", err.Error()))
	}
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, db db.DB, form *form.User) (*model.User, error) {

	user, err := s.UserRepository.UpdateUser(ctx, db, form)
	if err != nil {
		return &model.User{}, errors.New(fmt.Sprintf("current error %v ", err.Error()))
	}
	return user, nil
}

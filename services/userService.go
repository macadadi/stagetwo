package services

import (
	"context"

	"macadadi/stagetwo/db"
	"macadadi/stagetwo/form"
	"macadadi/stagetwo/model"
	"macadadi/stagetwo/repository"
)

type(
	UserServiceInterface interface{
		AddUser(ctx context.Context, db db.DB, form *form.User)error
		GetAllUsers(ctx context.Context, db db.DB)([]*model.User,error)
	}

	UserService struct{
		UserRepository *repository.UserRepository
	}
)

func NewUserService( repo *repository.UserRepository)*UserService{
	return &UserService{
		UserRepository: repo,
	}
}

func(s *UserService)AddUser(ctx context.Context, db db.DB, form *form.User)(error){	
	err := s.UserRepository.AddUser(ctx, db, form)

	if err != nil{
		return err
	}
	return nil
}

func(s *UserService)GetAllUsers(ctx context.Context, db db.DB)([]*model.User,error){
	users, err := s.UserRepository.GetAllUsers(ctx,db)

	if err != nil{
		return []*model.User{},err
	}
	return users,nil
}
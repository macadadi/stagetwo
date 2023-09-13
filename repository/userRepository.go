package repository

import (
	"context"

	"macadadi/stagetwo/db"
	"macadadi/stagetwo/form"
	"macadadi/stagetwo/model"
)

const(
        Adduser = "INSERT INTO users (name) VALUES ($1)"
		getAllUsers = "SELECT id, name  FROM users"
)

type(
	UserInterface interface{
		AddUser(ctx context.Context, db db.DB, form *form.User)error
		GetAllUsers(ctx context.Context, db db.DB)([]*model.User, error)
	}
	 UserRepository struct{}
)

func NewUserRepository()*UserRepository{
	return &UserRepository{}
}

func(u *UserRepository)AddUser(ctx context.Context, db db.DB, form *form.User)error{
	  _, err := db.ExecContext(ctx,Adduser,&form.Name)
	  if err != nil{
		return err
	  }
	  return nil
}

func (s *UserRepository)GetAllUsers(ctx context.Context, db db.DB)([]*model.User, error){
	rows, err := db.QueryContext(ctx, getAllUsers)

	if err != nil{
		return []*model.User{}, err
	}
	var users = make([]*model.User,0)
	for rows.Next(){
		var user model.User

		 rows.Scan(	
			&user.Id,
			&user.Full_name,
			&user.Country_code,		
			&user.Created_at,
		
		)
		if err != nil{
			return []*model.User{},err
		}
	    users = append(users, &user)
	}
	return users, nil
}
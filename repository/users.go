package repository

import (
	"context"

	"baptiste.com/models"
)

type UsersRepository interface {
	InsertUser(ctx context.Context, user *models.UserInsert) error
	GetUser(ctx context.Context, id string) (*models.Users, error)
	GetAllUsers(ctx context.Context) (*[]models.Users, error)
	UpdateUser(ctx context.Context, user *models.Users) error
}

var usersImplementation UsersRepository

func SetUsersRepository(repository UsersRepository) {
	usersImplementation = repository
}

func InsertUser(ctx context.Context, user *models.UserInsert) error {
	return usersImplementation.InsertUser(ctx, user)
}

func GetUser(ctx context.Context, id string) (*models.Users, error) {
	return usersImplementation.GetUser(ctx, id)
}

func GetAllUsers(ctx context.Context) (*[]models.Users, error) {
	return usersImplementation.GetAllUsers(ctx)
}

func UpdateUser(ctx context.Context, user *models.Users) error {
	return usersImplementation.UpdateUser(ctx, user)
}

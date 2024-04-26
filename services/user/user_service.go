package user

import (
	"context"
	db "expense-manager-backend/db/sqlc"
)

type UserService struct {
	db *db.Queries
}

func NewUserService(db *db.Queries) *UserService {
	return &UserService{
		db: db,
	}
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

func (s *UserService) CreateUser(ctx context.Context, params CreateUserParams) (db.User, error) {
	hashedPassword, e := GenerateFromPassword(params.Password)
	if e != nil {
		return db.User{}, e
	}

	newUser, e := s.db.CreateUser(ctx, db.CreateUserParams{
		Name:     params.Name,
		Email:    params.Email,
		Password: hashedPassword,
		ID:       db.GenerateId(),
	})

	return newUser, e
}

func (s *UserService) GetUserFromId(ctx context.Context, id string) (db.User, error) {
	existingUser, e := s.db.GetUser(ctx, id)

	return existingUser, e
}

func (s *UserService) GetUserFromEmail(ctx context.Context, email string) (db.User, error) {
	existingUser, e := s.db.GetUserFromEmail(ctx, email)

	return existingUser, e
}

type UpdateUserParams struct {
	Name  string
	Email string
}

func (s *UserService) UpdateUser(ctx context.Context, id string, params UpdateUserParams) (db.User, error) {
	updatedUser, e := s.db.UpdateUser(ctx, db.UpdateUserParams{
		ID:    id,
		Name:  params.Name,
		Email: params.Email,
	})

	return updatedUser, e
}

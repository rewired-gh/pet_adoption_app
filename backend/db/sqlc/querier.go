// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package sqlc

import (
	"context"
)

type Querier interface {
	AddContact(ctx context.Context, arg AddContactParams) error
	AuthUser(ctx context.Context, arg AuthUserParams) (int32, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) error
	CreatePet(ctx context.Context, arg CreatePetParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	GetCategoryID(ctx context.Context, arg GetCategoryIDParams) (int32, error)
	GetRoles(ctx context.Context, uid int32) ([]string, error)
	GetUser(ctx context.Context, uid int32) (GetUserRow, error)
	ListPet(ctx context.Context) ([]ListPetRow, error)
}

var _ Querier = (*Queries)(nil)

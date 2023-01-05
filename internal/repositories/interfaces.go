package repositories

import (
	"context"
)

type UserRepository interface {
	CreateUser(context.Context, string) (int, error)
}

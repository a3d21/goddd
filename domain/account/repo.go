package account

import (
	"context"
)

type Repo interface {
	CreateAccount(ctx context.Context, name string) (User, error)
	DeleteAccount(ctx context.Context, user User) (Nothing, error)
	AddAmount(ctx context.Context, user User, amount Amount) (Amount, error)
	MinusAmount(ctx context.Context, user User, amount Amount) (Amount, error)
	GetAmount(ctx context.Context, user User) (Amount, error)
}

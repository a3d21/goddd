package meminfra

import (
	"context"
	"fmt"

	"github.com/a3d21/goddd/base/baseerr"
	"github.com/a3d21/goddd/domain/account"
)

func NewAccountRepo() account.Repo {
	return &accountRepoImpl{data: map[account.User]account.Amount{}}
}

type accountRepoImpl struct {
	data map[account.User]account.Amount
}

func (r *accountRepoImpl) CreateAccount(ctx context.Context, name string) (account.User, error) {
	if _, exist := r.data[account.User(name)]; exist {
		return "", fmt.Errorf("duplicated name: %s. %w", name, baseerr.DuplicatedErr)
	}

	a := account.User(name)
	r.data[a] = 0
	return a, nil
}

func (r *accountRepoImpl) DeleteAccount(ctx context.Context, user account.User) (account.Nothing, error) {
	delete(r.data, user)
	return account.Nothing{}, nil

}

func (r *accountRepoImpl) AddAmount(ctx context.Context, user account.User, amount account.Amount) (account.Amount, error) {
	r.data[user] += amount
	return r.data[user], nil
}

func (r *accountRepoImpl) MinusAmount(ctx context.Context, user account.User, amount account.Amount) (account.Amount, error) {
	r.data[user] -= amount
	return r.data[user], nil
}

func (r *accountRepoImpl) GetAmount(ctx context.Context, user account.User) (account.Amount, error) {
	a, ok := r.data[user]
	if !ok {
		return 0, fmt.Errorf("user: %v not found. %w", user, baseerr.NotFoundErr)
	}
	return a, nil
}

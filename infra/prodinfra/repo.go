package prodinfra

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/a3d21/goddd/base/baseerr"
	"github.com/a3d21/goddd/domain/account"
)

func NewAccountRepo(db *gorm.DB) account.Repo {
	return &accountRepoImpl{db: db}
}

type accountRepoImpl struct {
	db *gorm.DB
}

func (r *accountRepoImpl) CreateAccount(ctx context.Context, name string) (account.User, error) {
	now := time.Now()
	po := &AccountPO{
		ID:        0,
		Name:      name,
		Amount:    0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := r.db.Save(po).Error
	if err != nil {
		if Duplicated(err) {
			return "", fmt.Errorf("%s %w", err.Error(), baseerr.DuplicatedErr)
		}
		return "", err
	}
	return account.User(name), nil
}

func (r *accountRepoImpl) DeleteAccount(ctx context.Context, user account.User) (account.Nothing, error) {
	err := r.db.Delete(&AccountPO{}, "name = ?", string(user)).Error
	return account.Nothing{}, err
}

func (r *accountRepoImpl) AddAmount(ctx context.Context, user account.User, amount account.Amount) (account.Amount, error) {
	acc, err := r.getAccount(ctx, user)
	if err != nil {
		return 0, err
	}
	acc.Amount += int64(amount)
	acc.UpdatedAt = time.Now()

	err = r.db.Save(acc).Error
	return account.Amount(acc.Amount), err
}

func (r *accountRepoImpl) MinusAmount(ctx context.Context, user account.User, amount account.Amount) (account.Amount, error) {
	acc, err := r.getAccount(ctx, user)
	if err != nil {
		return 0, err
	}
	acc.Amount -= int64(amount)
	acc.UpdatedAt = time.Now()

	err = r.db.Save(acc).Error
	return account.Amount(acc.Amount), err
}

func (r *accountRepoImpl) GetAmount(ctx context.Context, user account.User) (account.Amount, error) {
	acc, err := r.getAccount(ctx, user)
	if err != nil {
		return 0, err
	}

	return account.Amount(acc.Amount), nil
}

func (r *accountRepoImpl) getAccount(ctx context.Context, user account.User) (*AccountPO, error) {
	a := &AccountPO{}
	err := r.db.First(a, "name = ?", string(user)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s %w", err.Error(), baseerr.NotFoundErr)
		}
		return nil, err
	}
	return a, nil

}

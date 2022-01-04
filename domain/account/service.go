package account

import (
	"context"
)

type Service interface {
	// Open 开户
	Open(ctx context.Context, name string) (User, error)
	// Close 关闭账户
	Close(ctx context.Context, user User) (Nothing, error)
	// Credit 存钱
	Credit(ctx context.Context, user User, amount Amount) (Amount, error)
	// Debit 取钱
	Debit(ctx context.Context, user User, amount Amount) (Amount, error)
	// Balance 余额
	Balance(ctx context.Context, user User) (Amount, error)
}

func NewService(r Repo) Service {
	return &serviceImpl{r: r}
}

type serviceImpl struct {
	r Repo
}

func (s *serviceImpl) Open(ctx context.Context, name string) (User, error) {
	u, err := s.r.CreateAccount(ctx, name)
	return u, err
}

func (s *serviceImpl) Close(ctx context.Context, user User) (Nothing, error) {
	return s.r.DeleteAccount(ctx, user)
}

func (s *serviceImpl) Credit(ctx context.Context, user User, amount Amount) (Amount, error) {
	// TODO: valid user and perm first
	a, err := s.r.AddAmount(ctx, user, amount)
	if err != nil {
		// do something
		return 0, err
	}
	return a, err
}

func (s *serviceImpl) Debit(ctx context.Context, user User, amount Amount) (Amount, error) {
	return s.r.MinusAmount(ctx, user, amount)
}

func (s *serviceImpl) Balance(ctx context.Context, user User) (Amount, error) {
	a, err := s.r.GetAmount(ctx, user)
	if err != nil {
		// do something
		return 0, err
	}
	return a, nil
}

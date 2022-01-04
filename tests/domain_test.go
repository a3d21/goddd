package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"testing/quick"

	"github.com/a3d21/goddd/base/baseerr"
	"github.com/a3d21/goddd/di"
	"github.com/a3d21/goddd/domain/account"
)

var c = di.GetContainerByEnv()

// TestOpenThenClose 开户后关闭，查询余额应该返回 NotFound
func TestOpenThenClose(t *testing.T) {
	name := "6uklgyArAX0102XK"
	ctx := context.Background()

	e := c.Invoke(func(s account.Service) error {
		u, err := s.Open(ctx, name)
		if err != nil {
			return err
		}

		_, err = s.Close(ctx, u)
		if err != nil {
			return err
		}

		_, err = s.Balance(ctx, u)

		if !errors.Is(err, baseerr.NotFoundErr) {
			return errors.New("user should not found")
		}
		return nil
	})
	if e != nil {
		t.Error(e)
	}

}

// TestCreditThenDebit 存钱取钱相同金额，余额不变
func TestCreditThenDebit(t *testing.T) {
	name := "sD1ZePXKtDlMDCiy"
	someAmount := account.Amount(333)
	ctx := context.Background()
	e := c.Invoke(func(s account.Service) error {
		user, err := s.Open(ctx, name)
		if err != nil {
			return err
		}

		_, err = s.Credit(ctx, user, someAmount)
		if err != nil {
			return err
		}

		_, err = s.Debit(ctx, user, someAmount)
		if err != nil {
			return err
		}

		got, err := s.Balance(ctx, user)
		if err != nil {
			return err
		}

		if got != account.Amount(0) {
			return fmt.Errorf("%v != %v", account.Amount(0), got)
		}

		return nil
	})

	if e != nil {
		t.Error(e)
	}
}

// TestCreditAndDebitSpec 存、取款的属性(Property Testing)
// 任意存，并取相同数额的钱，帐号金额不变
func TestCreditAndDebitSpec(t *testing.T) {

	var s account.Service
	// get svc from container
	e := c.Invoke(func(svc account.Service) {
		s = svc
	})
	if e != nil {
		t.Error(e)
	}

	name := "hk1yboyiAMUymckO"
	ctx := context.Background()
	user, e := s.Open(ctx, name)
	if e != nil {
		t.Error(e)
	}

	assertion := func(seed uint) bool {
		amount := account.Amount(seed)

		balance, err := s.Balance(ctx, user)
		if err != nil {
			return false
		}

		_, err = s.Credit(ctx, user, amount)
		if err != nil {
			return false
		}

		_, err = s.Debit(ctx, user, amount)
		if err != nil {
			return false
		}

		got, err := s.Balance(ctx, user)
		if err != nil {
			return false
		}

		return got == balance

	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 2000,
	}); err != nil {
		t.Error(err)
	}
}

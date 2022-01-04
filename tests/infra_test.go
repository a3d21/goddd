package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/a3d21/goddd/base/baseerr"
	"github.com/a3d21/goddd/domain/account"
)

func TestDoubleCreateAccountShouldDuplicatedErr(t *testing.T) {
	var r account.Repo
	err := c.Invoke(func(repo account.Repo) {
		r = repo
	})
	if err != nil {
		t.Error(err)
	}

	name := "3o14Tu3dHyTo9Al6"
	ctx := context.Background()
	_, err = r.CreateAccount(ctx, name)
	if err != nil {
		t.Error(err)
	}

	_, err = r.CreateAccount(ctx, name)

	if !errors.Is(err, baseerr.DuplicatedErr) {
		t.Errorf("err:%v should be duplicated", err)
	}
}

func TestCreateThenDelete(t *testing.T) {
	var r account.Repo
	err := c.Invoke(func(repo account.Repo) {
		r = repo
	})
	if err != nil {
		t.Error(err)
	}

	name := "U9s8hxz8IGtFp1Ld"
	ctx := context.Background()
	a, err := r.CreateAccount(ctx, name)
	if err != nil {
		t.Error(err)
	}

	_, err = r.DeleteAccount(ctx, a)
	if err != nil {
		t.Error(err)
	}

	_, err = r.GetAmount(ctx, a)
	if !errors.Is(err, baseerr.NotFoundErr) {
		t.Error("user should not found")
	}

}

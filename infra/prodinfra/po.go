package prodinfra

import (
	"time"
)

type AccountPO struct {
	ID        int64  `gorm:"primaryKey,autoIncrement"`
	Name      string `gorm:"uniqueIndex"`
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*AccountPO) TableName() string {
	return "account"
}

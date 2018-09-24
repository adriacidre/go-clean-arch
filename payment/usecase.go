package payment

import (
	"context"

	model "github.com/adriacidre/go-clean-arch/models"
)

// Usecase payment usecase interface
type Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*model.Payment, string, error)
	GetByID(ctx context.Context, id int64) (*model.Payment, error)
	Update(ctx context.Context, p *model.Payment) (*model.Payment, error)
	GetByPaymentID(ctx context.Context, name string) (*model.Payment, error)
	Store(context.Context, *model.Payment) (*model.Payment, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

package payment

import (
	"context"

	"github.com/adriacidre/go-clean-arch/models"
)

// Repository repository interface to interact with payment model
type Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.Payment, error)
	GetByID(ctx context.Context, id int64) (*models.Payment, error)
	GetByPaymentID(ctx context.Context, title string) (*models.Payment, error)
	Update(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	Store(ctx context.Context, p *models.Payment) (int64, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

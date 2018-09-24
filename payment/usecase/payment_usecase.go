package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/adriacidre/go-clean-arch/models"
	"github.com/adriacidre/go-clean-arch/payment"
)

type paymentUsecase struct {
	repo           payment.Repository
	contextTimeout time.Duration
}

// NewPayment constructor for the payment use case.
func NewPayment(a payment.Repository, timeout time.Duration) payment.Usecase {
	return &paymentUsecase{
		repo:           a,
		contextTimeout: timeout,
	}
}

// Fetch fetches a list of rows from the database from "cursor" and a limit of "num".
func (a *paymentUsecase) Fetch(c context.Context, cursor string, num int64) ([]*models.Payment, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPayment, err := a.repo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	if size := len(listPayment); size == int(num) {
		lastID := listPayment[num-1].ID
		nextCursor = strconv.Itoa(int(lastID))
	}

	return listPayment, nextCursor, nil
}

// GetByID get a payment by ID.
func (a *paymentUsecase) GetByID(c context.Context, id int64) (*models.Payment, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update update the given payment.
func (a *paymentUsecase) Update(c context.Context, ar *models.Payment) (*models.Payment, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.repo.Update(ctx, ar)
}

// GetByPaymentID get a payment by its name.
func (a *paymentUsecase) GetByPaymentID(c context.Context, name string) (*models.Payment, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err := a.repo.GetByPaymentID(ctx, name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Store stores the given payment on the repository.
func (a *paymentUsecase) Store(c context.Context, m *models.Payment) (*models.Payment, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	if existedPayment, _ := a.GetByPaymentID(ctx, m.PaymentID); existedPayment != nil {
		return nil, models.ErrConflict
	}

	id, err := a.repo.Store(ctx, m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

// Delete removes a payment by id on the repository.
func (a *paymentUsecase) Delete(c context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	existedPayment, _ := a.repo.GetByID(ctx, id)
	if existedPayment == nil {
		return false, models.ErrNotFound
	}

	return a.repo.Delete(ctx, id)
}

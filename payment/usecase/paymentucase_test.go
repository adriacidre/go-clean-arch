package usecase_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	models "github.com/adriacidre/go-clean-arch/models"
	"github.com/adriacidre/go-clean-arch/payment/mocks"
	ucase "github.com/adriacidre/go-clean-arch/payment/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockPaymentRepo := new(mocks.Repository)
	mockPayment := &models.Payment{
		PaymentID:    "Hello",
		Organisation: "Content",
	}

	mockListArtilce := make([]*models.Payment, 0)
	mockListArtilce = append(mockListArtilce, mockPayment)
	mockPaymentRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	u := ucase.NewPayment(mockPaymentRepo, time.Second*2)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
	cursorExpected := strconv.Itoa(int(mockPayment.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))

	mockPaymentRepo.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockPaymentRepo := new(mocks.Repository)

	mockPaymentRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpexted Error"))

	u := ucase.NewPayment(mockPaymentRepo, time.Second*2)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)

	assert.Empty(t, nextCursor)
	assert.Error(t, err)
	assert.Len(t, list, 0)
	mockPaymentRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockPaymentRepo := new(mocks.Repository)
	mockPayment := models.Payment{
		PaymentID:    "Hello",
		Organisation: "Content",
	}

	mockPaymentRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockPayment, nil)

	u := ucase.NewPayment(mockPaymentRepo, time.Second*2)

	a, err := u.GetByID(context.TODO(), mockPayment.ID)

	assert.NoError(t, err)
	assert.NotNil(t, a)

	mockPaymentRepo.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockPaymentRepo := new(mocks.Repository)
	mockPayment := models.Payment{
		PaymentID:    "Hello",
		Organisation: "Content",
	}
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockPayment := mockPayment
	tempMockPayment.ID = 0

	mockPaymentRepo.On("GetByPaymentID", mock.Anything, mock.AnythingOfType("string")).Return(nil, models.ErrNotFound)
	mockPaymentRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Payment")).Return(mockPayment.ID, nil)

	u := ucase.NewPayment(mockPaymentRepo, time.Second*2)

	a, err := u.Store(context.TODO(), &tempMockPayment)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockPayment.PaymentID, tempMockPayment.PaymentID)
	mockPaymentRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockPaymentRepo := new(mocks.Repository)
	mockPayment := models.Payment{
		PaymentID:    "Hello",
		Organisation: "Content",
	}

	mockPaymentRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockPayment, models.ErrNotFound)

	mockPaymentRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(true, nil)

	u := ucase.NewPayment(mockPaymentRepo, time.Second*2)

	a, err := u.Delete(context.TODO(), mockPayment.ID)

	assert.NoError(t, err)
	assert.True(t, a)
	mockPaymentRepo.AssertExpectations(t)
}

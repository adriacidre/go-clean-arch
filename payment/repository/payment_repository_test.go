package repository_test

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	models "github.com/adriacidre/go-clean-arch/models"
	paymentRepo "github.com/adriacidre/go-clean-arch/payment/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "payment_id", "organisation_id", "updated_at", "created_at"}).
		AddRow(1, "payment 1", "Organisation 1", time.Now(), time.Now()).
		AddRow(2, "payment 2", "Organisation 2", time.Now(), time.Now())

	query := "SELECT id,payment_id,organisation, updated_at, created_at FROM payment WHERE ID > \\? LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := paymentRepo.NewMysqlPayment(db)
	cursor := "sampleCursor"
	num := int64(5)
	list, err := a.Fetch(context.TODO(), cursor, num)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "payment_id", "organisation_id", "updated_at", "created_at"}).
		AddRow(1, "payment 1", "Organisation 1", time.Now(), time.Now())

	query := "SELECT id,payment_id,organisation, updated_at, created_at FROM payment WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := paymentRepo.NewMysqlPayment(db)

	num := int64(5)
	anPayment, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anPayment)
}

func TestStore(t *testing.T) {
	now := time.Now()
	ar := &models.Payment{
		PaymentID:    "Judul",
		Organisation: "Organisation",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT  payment SET payment_id=\\? , organisation=\\? , updated_at=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.PaymentID, ar.Organisation, AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(12, 1))

	a := paymentRepo.NewMysqlPayment(db)

	lastID, err := a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), lastID)
}

func TestGetByPaymentID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "payment_id", "organisation_id", "updated_at", "created_at"}).
		AddRow(1, "payment 1", "Organisation 1", time.Now(), time.Now())

	query := "SELECT id,payment_id,organisation, updated_at, created_at FROM payment WHERE payment_id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := paymentRepo.NewMysqlPayment(db)

	payment := "payment 1"
	anPayment, err := a.GetByPaymentID(context.TODO(), payment)
	assert.NoError(t, err)
	assert.NotNil(t, anPayment)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "DELETE FROM payment WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := paymentRepo.NewMysqlPayment(db)

	num := int64(12)
	anPaymentStatus, err := a.Delete(context.TODO(), num)
	assert.NoError(t, err)
	assert.True(t, anPaymentStatus)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	ar := &models.Payment{
		ID:           12,
		PaymentID:    "Judul",
		Organisation: "Organisation",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "UPDATE payment set payment_id=\\?, organisation=\\?, updated_at=\\? WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.PaymentID, ar.Organisation, AnyTime{}, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := paymentRepo.NewMysqlPayment(db)

	s, err := a.Update(context.TODO(), ar)
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

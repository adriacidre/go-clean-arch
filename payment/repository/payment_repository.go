package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	models "github.com/adriacidre/go-clean-arch/models"
	payment "github.com/adriacidre/go-clean-arch/payment"
)

type mysqlPayment struct {
	Conn *sql.DB
}

// NewMysqlPayment mysql payment constructor.
func NewMysqlPayment(Conn *sql.DB) payment.Repository {
	return &mysqlPayment{Conn}
}

func (m *mysqlPayment) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Payment, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Payment, 0)
	for rows.Next() {
		t := new(models.Payment)
		err = rows.Scan(
			&t.ID,
			&t.PaymentID,
			&t.Organisation,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlPayment) Fetch(ctx context.Context, cursor string, num int64) ([]*models.Payment, error) {
	query := `SELECT id,payment_id,organisation, updated_at, created_at
  						FROM payment WHERE ID > ? LIMIT ?`

	return m.fetch(ctx, query, cursor, num)
}

func (m *mysqlPayment) GetByID(ctx context.Context, id int64) (a *models.Payment, err error) {
	query := `SELECT id,payment_id,organisation, updated_at, created_at
  						FROM payment WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (m *mysqlPayment) GetByPaymentID(ctx context.Context, payment string) (a *models.Payment, err error) {
	query := `SELECT id,payment_id,organisation, updated_at, created_at
  						FROM payment WHERE payment_id = ?`

	list, err := m.fetch(ctx, query, payment)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (m *mysqlPayment) Store(ctx context.Context, a *models.Payment) (int64, error) {
	query := `INSERT payment SET payment_id=? , organisation=? , updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {

		return 0, err
	}

	logrus.Debug("Created At: ", a.CreatedAt)
	res, err := stmt.ExecContext(ctx, a.PaymentID, a.Organisation, time.Now(), time.Now())
	if err != nil {

		return 0, err
	}
	return res.LastInsertId()
}

func (m *mysqlPayment) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM payment WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {

		return false, err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		logrus.Error(err)
		return false, err
	}

	return true, nil
}

func (m *mysqlPayment) Update(ctx context.Context, ar *models.Payment) (*models.Payment, error) {
	query := `UPDATE payment set payment_id=?, organisation=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, ar.PaymentID, ar.Organisation, time.Now(), ar.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		logrus.Error(err)
		return nil, err
	}

	return ar, nil
}

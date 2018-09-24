package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	models "github.com/adriacidre/go-clean-arch/models"
	paymentHttp "github.com/adriacidre/go-clean-arch/payment/delivery/http"
	"github.com/adriacidre/go-clean-arch/payment/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/faker"
)

func TestFetch(t *testing.T) {
	var mockPayment models.Payment
	err := faker.FakeData(&mockPayment)
	assert.NoError(t, err)
	mockUCase := new(mocks.Payment)
	mockListPayment := make([]*models.Payment, 0)
	mockListPayment = append(mockListPayment, &mockPayment)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(mockListPayment, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/payment?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := paymentHttp.PaymentHandler{
		Usecase: mockUCase,
	}
	assert.Nil(t, handler.FetchPayment(c))

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.Payment)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(nil, "", models.ErrInternalServer)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/payment?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := paymentHttp.PaymentHandler{
		Usecase: mockUCase,
	}
	assert.Nil(t, handler.FetchPayment(c))

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockPayment models.Payment
	err := faker.FakeData(&mockPayment)
	assert.NoError(t, err)

	mockUCase := new(mocks.Payment)

	num := int(mockPayment.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(&mockPayment, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/payment/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("payment/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := paymentHttp.PaymentHandler{
		Usecase: mockUCase,
	}
	assert.Nil(t, handler.GetByID(c))

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockPayment := models.Payment{
		PaymentID:    "Payment",
		Organisation: "ORG",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	tempMockPayment := mockPayment
	tempMockPayment.ID = 0
	mockUCase := new(mocks.Payment)

	j, err := json.Marshal(tempMockPayment)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*models.Payment")).Return(&mockPayment, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/payment", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/payment")

	handler := paymentHttp.PaymentHandler{
		Usecase: mockUCase,
	}
	assert.Nil(t, handler.Store(c))

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	var mockPayment models.Payment
	err := faker.FakeData(&mockPayment)
	assert.NoError(t, err)

	mockUCase := new(mocks.Payment)

	num := int(mockPayment.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(&mockPayment, nil)
	mockUCase.On("Update", mock.Anything, mock.AnythingOfType("*models.Payment")).Return(&mockPayment, nil)

	tempMockPayment := mockPayment
	tempMockPayment.Organisation = "modified"

	j, err := json.Marshal(tempMockPayment)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.PATCH, "/payment/"+strconv.Itoa(num), strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("payment/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := paymentHttp.PaymentHandler{
		Usecase: mockUCase,
	}
	assert.Nil(t, handler.Update(c))

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockPayment models.Payment
	err := faker.FakeData(&mockPayment)
	assert.NoError(t, err)

	mockUCase := new(mocks.Payment)

	num := int(mockPayment.ID)

	mockUCase.On("Delete", mock.Anything, int64(num)).Return(true, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/payment/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("payment/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := paymentHttp.PaymentHandler{
		Usecase: mockUCase,
	}
	assert.Nil(t, handler.Delete(c))

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)

}

package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"

	models "github.com/adriacidre/go-clean-arch/models"

	paymentUcase "github.com/adriacidre/go-clean-arch/payment"
	"github.com/labstack/echo"

	validator "gopkg.in/go-playground/validator.v9"
)

// ResponseError response struct representing an error.
type ResponseError struct {
	Message string `json:"message"`
}

// PaymentHandler http handler for payment use cases.
type PaymentHandler struct {
	Usecase paymentUcase.Usecase
}

// NewPaymentHTTPHandler payment http handler constructor.
func NewPaymentHTTPHandler(e *echo.Echo, us paymentUcase.Usecase) {
	handler := &PaymentHandler{
		Usecase: us,
	}
	e.GET("/payment", handler.FetchPayment)
	e.POST("/payment", handler.Store)
	e.PATCH("/payment/:id", handler.Update)
	e.GET("/payment/:id", handler.GetByID)
	e.DELETE("/payment/:id", handler.Delete)
}

// FetchPayment handles fetching lists of payments.
func (h *PaymentHandler) FetchPayment(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, nextCursor, err := h.Usecase.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)

	return c.JSON(http.StatusOK, listAr)
}

// GetByID handles geting payments by ID requests.
func (h *PaymentHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Input ID is not valid"})
	}
	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := h.Usecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

// isRequestValid validates request mapped payment.
func isRequestValid(m *models.Payment) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Store handles the payment storage requests.
func (h *PaymentHandler) Store(c echo.Context) error {
	var payment models.Payment

	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&payment); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := h.Usecase.Store(ctx, &payment)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, ar)
}

// Delete handler payment removal requests.
func (h *PaymentHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Input ID is not valid"})
	}
	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if _, err = h.Usecase.Delete(ctx, id); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// Update handles the payment storage updates.
func (h *PaymentHandler) Update(c echo.Context) error {
	var input models.Payment

	if err := c.Bind(&input); err != nil {
		spew.Dump(input)
		println(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&input); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Input ID is not valid"})
	}
	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	payment, err := h.Usecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	payment.Organisation = input.Organisation
	ar, err := h.Usecase.Update(ctx, payment)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ar)
}

// getStatusCode based on the useacse output error calculates the http response
// status code.
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.ErrInternalServer:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

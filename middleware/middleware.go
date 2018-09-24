package middleware

import "github.com/labstack/echo"

const (
	// AccessTokenKey valid access token key.
	AccessTokenKey = "Access-Token"
)

type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// InitMiddleware initializes middleware.
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

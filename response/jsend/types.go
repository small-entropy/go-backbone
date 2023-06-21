package jsend

import "github.com/labstack/echo/v4"

type Response struct {
	Status  string    `json:"status" yaml:"status"`
	Data    *echo.Map `json:"data,omitempty" yaml:"data,omitempty"`
	Message string    `json:"message,omitempty" yaml:"message,omitempty"`
	Code    int       `json:"code,omitempty" yaml:"code,omitempty"`
}

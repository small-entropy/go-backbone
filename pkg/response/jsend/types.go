package jsend

import (
	facade "github.com/small-entropy/go-backbone/third_party/facade/echo"
)

type Response struct {
	Status  string      `json:"status" yaml:"status"`
	Data    *facade.Map `json:"data,omitempty" yaml:"data,omitempty"`
	Message string      `json:"message,omitempty" yaml:"message,omitempty"`
	Code    int         `json:"code,omitempty" yaml:"code,omitempty"`
}

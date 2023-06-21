package server

import (
	"github.com/small-entropy/go-backbone/database/abstract"

	"github.com/labstack/echo/v4"
)

type RoutesOptions struct {
	Path    string
	Method  string
	Handler echo.HandlerFunc
}

type ServerSettings struct {
	Address       string
	RoutesOptions []RoutesOptions
	IndexOptions  []abstract.IndexOptions
}

type Server struct {
	Core     *echo.Echo
	Settings ServerSettings
}

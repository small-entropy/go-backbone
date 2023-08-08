package server

import (
	"github.com/small-entropy/go-backbone/database/abstract"

	"github.com/labstack/echo/v4"
)

type Echo = echo.Echo

type EchoHandlerFunc = echo.HandlerFunc

type RoutesOptions struct {
	Path    string
	Method  string
	Handler EchoHandlerFunc
}

type ServerSettings struct {
	Address       string
	RoutesOptions []RoutesOptions
	IndexOptions  []abstract.IndexOptions
}

type Server struct {
	Core     Echo
	Settings ServerSettings
}

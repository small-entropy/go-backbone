package server

import (
	"github.com/small-entropy/go-backbone/database/abstract"

	echo_facade "github.com/small-entropy/go-backbone/facades/echo"
)

type RoutesOptions struct {
	Path    string
	Method  string
	Handler echo_facade.EchoHandlerFunc
}

type ServerSettings struct {
	Address       string
	RoutesOptions []RoutesOptions
	IndexOptions  []abstract.IndexOptions
}

type Server struct {
	Core     echo_facade.Echo
	Settings ServerSettings
}

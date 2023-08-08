package server

import (
	"github.com/small-entropy/go-backbone/pkg/database/abstract"

	echo_facade "github.com/small-entropy/go-backbone/third_party/facade/echo"
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

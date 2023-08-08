package server

import (
	"github.com/small-entropy/go-backbone/pkg/server/abstract"

	facade "github.com/small-entropy/go-backbone/third_party/facade/echo"

	"errors"
)

type RoutesOptions = abstract.RoutesOptions[facade.EchoHandlerFunc]

type ServerSettings = abstract.ServerSettings[facade.EchoHandlerFunc]

type EchoServer struct {
	abstract.Server[*facade.Echo, facade.EchoHandlerFunc]
}

func (s *EchoServer) GetCore() *facade.Echo {
	return s.Server.Core
}

func (s *EchoServer) RegisterRoutes() error {
	var err error
	for _, v := range s.Settings.RoutesOptions {
		if err = s.RegisterRoute(&v); err != nil {
			break
		}
	}
	return err
}

func (s *EchoServer) RegisterRoute(opt *RoutesOptions) error {
	var err error
	switch opt.Method {
	case "POST":
		s.Core.POST(opt.Path, opt.Handler)
	case "GET":
		s.Core.GET(opt.Path, opt.Handler)
	case "PUT":
		s.Core.PUT(opt.Path, opt.Handler)
	case "DELETE":
		s.Core.DELETE(opt.Path, opt.Handler)
	default:
		err = errors.New("wrong data")
	}
	return err
}

func (s *EchoServer) Run() error {
	return s.Core.Start(s.Settings.Address)
}

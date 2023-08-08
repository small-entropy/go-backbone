package server

import "errors"

func (s *Server) RegisterRoutes() error {
	var err error
	for _, v := range s.Settings.RoutesOptions {
		if err = s.RegisterRoute(&v); err != nil {
			break
		}
	}
	return err
}

func (s *Server) RegisterRoute(opt *RoutesOptions) error {
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

func (s *Server) Run() error {
	return s.Core.Start(s.Settings.Address)
}

package abstract

import "github.com/small-entropy/go-backbone/pkg/database/abstract"

type RoutesOptions[H any] struct {
	Path    string
	Method  string
	Handler H
}

type ServerSettings[H any] struct {
	Address       string
	RoutesOptions []RoutesOptions[H]
	IndexOptions  []abstract.IndexOptions
}

type Server[C any, H any] struct {
	Core     C
	Settings ServerSettings[H]
}

type IServer[C any, H any] interface {
	GetCore() C
	RegisterRoutes() error
	RegisterRoute(opt *RoutesOptions[H]) error
	Run() error
}

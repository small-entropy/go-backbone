package abstract

type Database[C any, O any] struct {
	Name    string
	Client  C
	Options O
}

type IndexOptions struct {
	Collection string
	Name       string
	Unique     bool
}

package interfaces

type ISchema[I any, D any, T any] interface {
	GetIdentifier() I
	GetData() D
	Created() *T
	Updated() *T
	Deleted() *T
}

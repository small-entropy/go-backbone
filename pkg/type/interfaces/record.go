package interfaces

type IRecord interface {
	Data() map[string]interface{}
	Fields() []string
	Values() []interface{}
	HasField(field string) bool
	Value(field string) (interface{}, error)
	Add(field string, value interface{}) error
	Set(field string, value interface{}) error
	Version() int
	UpdateVersion()
}

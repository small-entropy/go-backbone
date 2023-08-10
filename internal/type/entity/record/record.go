package record

import (
	"errors"
)

type Record struct {
	fields map[int]string
	values map[int]interface{}
}

func (r *Record) Data() map[string]interface{} {
	data := map[string]interface{}{}
	for i, v := range r.fields {
		data[v] = r.values[i]
	}
	return data
}

func (r *Record) Fields() []string {
	var fields []string
	for _, v := range r.fields {
		fields = append(fields, v)
	}
	return fields
}

func (r *Record) Values() []interface{} {
	var values []interface{}
	for _, v := range r.values {
		values = append(values, v)
	}
	return values
}

func (r *Record) HasField(field string) bool {
	exist := false
	for _, v := range r.fields {
		if v == field {
			exist = true
			break
		}
	}
	return exist
}

func (r *Record) Value(field string) (interface{}, error) {
	var value interface{}
	var err error
	for i, v := range r.fields {
		if field == v {
			value = r.values[i]
			break
		}
	}

	if value == nil {
		err = errors.New("can not get value")
	}
	return value, err
}

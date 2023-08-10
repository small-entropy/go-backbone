package record

import (
	"errors"
)

type Record struct {
	fields  map[int]string
	values  map[int]interface{}
	version int
}

// New
// Конструктор Record
func New(fields *map[int]string, values *map[int]interface{}) *Record {
	return &Record{
		fields:  *fields,
		values:  *values,
		version: 0,
	}
}

func (r *Record) Version() int {
	return r.version
}

func (r *Record) UpdateVersion() {
	r.version += 1
}

// Data
// Метод получения всех данных
func (r *Record) Data() map[string]interface{} {
	data := map[string]interface{}{}
	for i, v := range r.fields {
		data[v] = r.values[i]
	}
	return data
}

// Fields
// Метод получения всех полей (ключей)
func (r *Record) Fields() []string {
	var fields []string
	for _, v := range r.fields {
		fields = append(fields, v)
	}
	return fields
}

// Values
// Метод получения всех значений
func (r *Record) Values() []interface{} {
	var values []interface{}
	for _, v := range r.values {
		values = append(values, v)
	}
	return values
}

// HasField
// Метод проверяет есть ли поле
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

// Value
// Метод получения значения по имени поля
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

// Add
// Метод добавления нового поля
func (r *Record) Add(field string, value interface{}) error {
	var err error
	hasField := r.HasField(field)
	if hasField != true {
		size := len(r.fields)
		newIndex := size + 1
		r.fields[newIndex] = field
		r.values[newIndex] = value
		r.UpdateVersion()
	} else {
		err = errors.New("property has exist. Can not add new property")
	}
	return err
}

// Set
// Метод обновления значения поля
func (r *Record) Set(field string, value interface{}) error {
	var err error
	hasField := r.HasField(field)
	if hasField == true {
		for i, v := range r.fields {
			if v == field {
				if r.values[i] != value {
					r.values[i] = value
					r.UpdateVersion()
				} else {
					err = errors.New("value are equals. Can not set new value")
				}
				break
			}
		}
	} else {
		err = errors.New("property not exist. Can not set new value")
	}
	return err
}

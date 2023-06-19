package gocsvparser

import (
	"encoding/csv"
)

type FieldsHandlerByName interface {
	FieldsName() string
	FieldsHandler
}

type FieldsHandlerByIndices interface {
	FieldsIndices() []int
	FieldByIndex(index int, field string) error
	FieldsHandler
}

type FieldsHandler interface {
	Fields(fields ...string) error
	NumFields() int
}

type RecordHandler interface {
	FieldsHandlers() []FieldsHandler
}

type Unmarshaler struct {
	header        []string
	reader        *csv.Reader
	recordHandler RecordHandler
}

func NewUnmarshaler() *Unmarshaler {
	return &Unmarshaler{}
}

func (u *Unmarshaler) WithCsvReader(reader *csv.Reader) *Unmarshaler {
	u.reader = reader
	return u
}

func (u *Unmarshaler) WithHeader(header []string) *Unmarshaler {
	if len(header) > 0 {
		newHeader := make([]string, len(header))
		copy(newHeader, header)
		u.header = newHeader
	}
	return u
}

func (u *Unmarshaler) WithRecordHandler(handler RecordHandler) *Unmarshaler {
	u.recordHandler = handler
	return u
}

func (u *Unmarshaler) Unmarshal(data []byte, v interface{}) error {
	//TODO: implementation
	return nil
}

func Unmarshal(data []byte, v interface{}) error {
	return NewUnmarshaler().Unmarshal(data, v)
}

// func Read(i interface{}) {
// 	var val reflect.Value

// 	val = reflect.ValueOf(i)
// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()

// 		if val.Kind() == reflect.Slice {
// 			val = val.Elem()

// 			if val.Kind() == reflect.Struct {
// 				fields := reflect.VisibleFields(val.Type())
// 			}
// 		}
// 	}
// }

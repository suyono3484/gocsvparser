package gocsvparser

import (
	"fmt"
	"reflect"
)

type defaultRecordHandlerMode int

const (
	single defaultRecordHandlerMode = 1
	slice  defaultRecordHandlerMode = 2
)

type defaultRecordHandler struct {
	structType     reflect.Type
	mode           defaultRecordHandlerMode
	slice          reflect.Value
	fieldByName    map[string]FieldsHandlerByName
	fieldByIndices map[int]FieldsHandlerByIndices
	fieldsHandlers []FieldsHandler
}

type defaultFieldsHandlerByName struct {
}

func newDefaultRecordHandler(v interface{}) (*defaultRecordHandler, error) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()

		recordHandler := &defaultRecordHandler{}
		switch val.Kind() {
		case reflect.Slice:
			recordHandler.mode = slice
			typ := val.Type().Elem()
			if typ.Kind() == reflect.Struct {
				recordHandler.slice = val
				recordHandler.structType = typ
			}
		case reflect.Struct:
			recordHandler.mode = single
		}
	}

	return nil, fmt.Errorf("invalid value %v", val)
}

func (dr *defaultRecordHandler) buildFieldsHandler() {

}

func (df *defaultFieldsHandlerByName) FieldsName() string {
	//TODO: implementation
	return ""
}

func (df *defaultFieldsHandlerByName) NumFields() int {
	//TODO: implementation
	return 0
}

func (df *defaultFieldsHandlerByName) Fields(fields ...string) error {
	//TODO: implementation
	return nil
}

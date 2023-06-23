package gocsvparser

import (
	"fmt"
	"reflect"
	"strings"
)

type defaultRecordHandler struct {
	handlersByName map[string]reflect.StructField
	outType        reflect.Type
	fieldsHandlers []FieldsHandler
}

func newDefaultHandler() *defaultRecordHandler {
	newHander := new(defaultRecordHandler)
	newHander.handlersByName = make(map[string]reflect.StructField)
	return newHander
}

func (d *defaultRecordHandler) HandleRecord(v interface{}, record []string) error {
	//TODO: implementation
	if d.outType == nil {

	}
	return nil
}

func (d *defaultRecordHandler) FieldsHandlers() []FieldsHandler {
	//TODO: implementation
	return nil
}

func (d *defaultRecordHandler) SetFieldConfigs(configs []FieldsConfig) {

}

func (d *defaultRecordHandler) parseVal(v interface{}) error {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()

		if typ.Kind() == reflect.Struct {
			d.outType = typ
			return d.buildStructHandlers()
		} else if typ.Kind() == reflect.Map {
			//TODO: implementation
		} else if typ.Kind() == reflect.Slice {
			//TODO: implementation
		}
	}
	return fmt.Errorf("v should be pointer of Struct, Map, or Slice: %+v", typ)
}

func (d *defaultRecordHandler) buildStructHandlers() error {
	//TODO: implementation
	for _, field := range reflect.VisibleFields(d.outType) {
		if csv, ok := field.Tag.Lookup(csvTag); ok {
			s := strings.Split(csv, ",")
			if len(s) == 0 {
				return fmt.Errorf("invalid tag %+v", field.Tag)
			}
			if _, ok = d.handlersByName[s[0]]; ok {
				return fmt.Errorf("problem with the receiving struct, multiple field with tag %s", s[0])
			}
			d.handlersByName[s[0]] = field
		} else if csvIndex, ok := field.Tag.Lookup(csvIndexTag); ok {
			_ = csvIndex //TODO: process tag
		}
	}
	return nil
}

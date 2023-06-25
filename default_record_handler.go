package gocsvparser

import (
	"fmt"
	"reflect"
	"strings"
)

type defaultRecordHandler struct {
	fieldByTag       map[string]columnFieldBinding
	outType          reflect.Type
	columnNameMapped bool
}

type columnFieldBinding struct {
	field reflect.StructField
	index int
}

func newDefaultHandler() *defaultRecordHandler {
	newHandler := new(defaultRecordHandler)
	newHandler.fieldByTag = make(map[string]columnFieldBinding)
	newHandler.columnNameMapped = false
	return newHandler
}

func (d *defaultRecordHandler) HandleRecord(v any, record []string) error {
	var (
		err error
	)

	if d.outType == nil {
		err = d.detType(v)
		if err != nil {
			return fmt.Errorf("error HandleRecord: detType: %+v", err)
		}

		err = d.mapStructTag()
		if err != nil {
			return fmt.Errorf("error HandleRecord: mapStructTag: %+v", err)
		}
	}

	if !d.columnNameMapped {

	}

	return nil
}

func (d *defaultRecordHandler) SetFieldConfigs(configs []FieldsConfig) {

}

func (d *defaultRecordHandler) detType(v any) error {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()

		if typ.Kind() == reflect.Struct {
			d.outType = typ
			return d.mapStructTag()
		} else if typ.Kind() == reflect.Map {
			//TODO: implementation
		} else if typ.Kind() == reflect.Slice {
			//TODO: implementation
		}
	}
	return fmt.Errorf("v should be pointer of Struct, Map, or Slice: %+v", typ)
}

func (d *defaultRecordHandler) mapStructTag() error {
	//TODO: implementation
	for _, field := range reflect.VisibleFields(d.outType) {
		if csv, ok := field.Tag.Lookup(csvTag); ok {
			s := strings.Split(csv, ",")
			if len(s) == 0 {
				return fmt.Errorf("invalid tag %+v", field.Tag)
			}
			if _, ok = d.fieldByTag[s[0]]; ok {
				return fmt.Errorf("problem with the receiving struct, multiple field with tag %s", s[0])
			}
			d.fieldByTag[s[0]] = columnFieldBinding{
				field: field,
			}
		} else if csvIndex, ok := field.Tag.Lookup(csvIndexTag); ok {
			_ = csvIndex //TODO: process tag
		}
	}
	return nil
}

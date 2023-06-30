package gocsvparser

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
		err      error
		strField string
		index    int
		ok       bool
		binding  columnFieldBinding
		val      reflect.Value
	)

	if d.outType == nil {
		err = d.detType(v)
		if err != nil {
			return fmt.Errorf("error HandleRecord: detType: %+v", err)
		}
	}

	if !d.columnNameMapped {
		for index, strField = range record {
			if binding, ok = d.fieldByTag[strField]; ok {
				binding.index = index
				d.fieldByTag[strField] = binding
			}
		}

		d.columnNameMapped = true
		return HeaderRead
	}

	val = reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer {
		return errors.New("error HandleRecord: v Kind() is not Pointer")
	}

	val = val.Elem()
	switch val.Kind() {
	case reflect.Struct:
		for _, binding = range d.fieldByTag {
			if binding.index < len(record) {
				val, err = d.setValue(val, record[binding.index], binding.field)
				if err != nil {
					return fmt.Errorf("error HandleRecord: StructField SetValue: %+v", err)
				}
			}
		}
	case reflect.Map:
		return errors.New("error HandleRecord: Map is not supported yet") //TODO: fix me
	case reflect.Slice:
		return errors.New("error HandleRecord: Slice is not supported yet") //TODO: fix me
	}

	return nil
}

func (d *defaultRecordHandler) setValue(val reflect.Value, strValue string, structFiled reflect.StructField) (reflect.Value, error) {
	var (
		fieldVal reflect.Value
		f64      float64
		i64      int64
		i        int
		err      error
		b        bool
	)

	fieldVal = val.FieldByIndex(structFiled.Index)
	switch fieldVal.Type().Kind() {
	case reflect.String:
		fieldVal.SetString(strValue)
	case reflect.Bool:
		b, err = strconv.ParseBool(strValue)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("ParseBool: %+v", err)
		}
		fieldVal.SetBool(b)
	case reflect.Int64:
		i64, err = strconv.ParseInt(strValue, 0, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("ParseInt 64: %+v", err)
		}
		fieldVal.SetInt(i64)
	case reflect.Int32:
		i64, err = strconv.ParseInt(strValue, 0, 32)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("ParseInt 32: %+v", err)
		}
		fieldVal.SetInt(i64)
	case reflect.Int16:
		i64, err = strconv.ParseInt(strValue, 0, 16)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("ParseInt 16: %+v", err)
		}
		fieldVal.SetInt(i64)
	case reflect.Int8:
		i64, err = strconv.ParseInt(strValue, 0, 8)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("ParseInt 8: %+v", err)
		}
		fieldVal.SetInt(i64)
	case reflect.Int:
		i, err = strconv.Atoi(strValue)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("strconv.Atoi: %+v", err)
		}
		fieldVal.SetInt(int64(i))
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		return reflect.Value{}, errors.New("unimplemented") //TODO: fix me
	case reflect.Float32:
		f64, err = strconv.ParseFloat(strValue, 32)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("ParseFloat 32: %+v", err)
		}
		fieldVal.SetFloat(f64)
	case reflect.Float64:
		f64, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("ParseFloat 64: %+v", err)
		}
		fieldVal.SetFloat(f64)
	default:
		return reflect.Value{}, errors.New("missing implementation") //TODO: fix me
	}

	return val, nil
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

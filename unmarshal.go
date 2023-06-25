package gocsvparser

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"
)

var (
	HeaderRead = errors.New("column headers successfully read")
)

type FieldsConfig struct {
	Name string
	Num  int
}

type FieldsHandlerByName interface {
	FieldName() string
	FieldByName(name, field string) error
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

type RecordFieldsHandler interface {
	FieldsHandlers() []FieldsHandler
	Out(v any) error
}

type RecordHandler interface {
	// SetFieldConfigs is only effective if a Map is passed to HandleRecord
	SetFieldConfigs(configs []FieldsConfig)
	HandleRecord(v any, record []string) error
}

type recordHandlerType int

const (
	direct         recordHandlerType = 1
	fieldsSpecific recordHandlerType = 2
)

type Unmarshaler struct {
	options            []CsvOption
	header             []string
	recordHandler      RecordHandler
	recordFieldHandler RecordFieldsHandler
	handlerType        recordHandlerType
	csvReader          *csv.Reader
}

func NewUnmarshaler(options ...CsvOption) *Unmarshaler {
	unmarshaler := &Unmarshaler{
		handlerType: direct,
		options:     []CsvOption{nil},
	}

	if len(options) > 0 {
		unmarshaler.options = append(unmarshaler.options, options...)
	}

	return unmarshaler
}

func (u *Unmarshaler) processOptions(options ...CsvOption) {
	for _, option := range options {
		if option == nil {
			continue
		}

		switch option.getType() {
		case csvReader:
			if u.csvReader != nil {
				continue
			}
			o := option.(csvReaderOption)
			u.csvReader = o.reader
		case comma:
			o := option.(commaOption)
			u.csvReader.Comma = o.comma
		case comment:
			o := option.(commentOption)
			u.csvReader.Comment = o.comment
		case fieldsPerRecord:
			o := option.(fieldsPerRecordOption)
			u.csvReader.FieldsPerRecord = o.fieldsPerRecord
		case lazyQuotes:
			o := option.(lazyQuotesOption)
			u.csvReader.LazyQuotes = o.lazyQuotes
		case trimLeadingSpace:
			o := option.(trimLeadingSpaceOption)
			u.csvReader.TrimLeadingSpace = o.trimLeadingSpace
		case reuseRecord:
			o := option.(reuseRecordOption)
			u.csvReader.ReuseRecord = o.reuseRecord
		case columnHeader:
			headerOption := option.(columnHeaderOption)
			u.header = headerOption.header
		case recordHandler:
			rho := option.(*recordHandlerOption)
			u.handlerType = rho.handlerType
			switch rho.handlerType {
			case direct:
				u.recordHandler = rho.recordHandler
			case fieldsSpecific:
				u.recordFieldHandler = rho.recordFieldsHandler
			}
		}
	}
}

func (u *Unmarshaler) Unmarshal(data []byte, v any, options ...CsvOption) error {
	var (
		typ        reflect.Type
		val, slice reflect.Value
		record     []string
		err        error
	)

	if len(options) > 0 {
		u.options = append(u.options, options...)
	}

	u.options[0] = CsvReader(csv.NewReader(bytes.NewReader(data)))
	u.processOptions(u.options...)

	if u.handlerType == direct {
		if u.recordHandler == nil {
			u.recordHandler = newDefaultHandler()
		}

		slice, typ, err = u.detOutputType(v)
		if err != nil {
			return fmt.Errorf("error Unmarshal: detOutputType: %+v", err)
		}

		for {
			record, err = u.csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("error Unmarshal: csv.Reader.Read: %+v", err)
			}

			val, err = u.newElem(typ)
			if err != nil {
				return fmt.Errorf("error Unmarshal: newElem: %+v", err)
			}

			err = u.recordHandler.HandleRecord(val.Interface(), record)
			if err != nil {
				if err == HeaderRead {
					continue
				}
				return fmt.Errorf("error Unmarshal: RecordHandler.HandleRecord: %+v", err)
			}

			slice.Set(reflect.Append(slice, val.Elem()))
		}
	} else {

	}

	//TODO: implementation
	return nil
}

func (u *Unmarshaler) detOutputType(v any) (reflect.Value, reflect.Type, error) {
	var typ reflect.Type

	if v == nil {
		return reflect.Value{}, nil, errors.New("the output parameter is nil")
	}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer {
		return reflect.Value{}, nil, errors.New("invalid output parameter type")
	}

	val = val.Elem()
	if val.Kind() != reflect.Slice {
		return reflect.Value{}, nil, errors.New("invalid output parameter type")
	}

	typ = val.Type().Elem()

	return val, typ, nil
}

func Unmarshal(data []byte, v any, options ...CsvOption) error {
	return NewUnmarshaler().Unmarshal(data, v, options...)
}

func (u *Unmarshaler) newElem(typ reflect.Type) (reflect.Value, error) {
	switch typ.Kind() {
	case reflect.Struct:
		return reflect.New(typ), nil
	case reflect.Map:
		//TODO: implementation
	case reflect.Slice:
		//TODO: implementation
	}

	return reflect.Zero(typ), errors.New("invalid impelementation") //TODO: placeholder; update me
}

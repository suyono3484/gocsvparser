package gocsvparser

import (
	"bytes"
	"encoding/csv"
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
	if len(options) > 0 {
		u.options = append(u.options, options...)
	}

	u.options[0] = CsvReader(csv.NewReader(bytes.NewReader(data)))
	u.processOptions(u.options...)

	if u.handlerType == direct {
		if u.recordHandler == nil {
			//TODO: build default generator
		}

	} else {

	}

	//TODO: implementation
	return nil
}

func Unmarshal(data []byte, v any, options ...CsvOption) error {
	return NewUnmarshaler().Unmarshal(data, v, options...)
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

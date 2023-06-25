package gocsvparser

import (
	"encoding/csv"
	"errors"
	"fmt"
)

type csvOptionsType int

type CsvOption interface {
	getType() csvOptionsType
}

const (
	comma csvOptionsType = iota
	comment
	fieldsPerRecord
	lazyQuotes
	trimLeadingSpace
	reuseRecord
	useCrlf
	columnHeader
	recordHandler
	csvReader
)

type csvReaderOption struct {
	reader *csv.Reader
}

func CsvReader(reader *csv.Reader) csvReaderOption {
	return csvReaderOption{
		reader: reader,
	}
}

func (csvReaderOption) getType() csvOptionsType {
	return csvReader
}

type recordHandlerOption struct {
	handlerType         recordHandlerType
	recordHandler       RecordHandler
	recordFieldsHandler RecordFieldsHandler
}

func RecordHandlerOption(handler any) (*recordHandlerOption, error) {
	var rho *recordHandlerOption
	if d, ok := handler.(RecordHandler); ok {
		if d == nil {
			return nil, errors.New("handler value is nil")
		}

		rho = &recordHandlerOption{
			handlerType:   direct,
			recordHandler: d,
		}
	} else if s, ok := handler.(RecordFieldsHandler); ok {
		if s == nil {
			return nil, errors.New("handler value is nil")
		}

		rho = &recordHandlerOption{
			handlerType:         fieldsSpecific,
			recordFieldsHandler: s,
		}
	} else {
		return nil, fmt.Errorf("invalid handler type %T", handler)
	}

	return rho, nil
}

func (r *recordHandlerOption) getType() csvOptionsType {
	return recordHandler
}

type columnHeaderOption struct {
	header []string
}

func ColumnHeader(header ...string) columnHeaderOption {
	columnHeader := columnHeaderOption{}
	if len(header) > 0 {
		columnHeader.header = make([]string, len(header))
		copy(columnHeader.header, header)
	}
	return columnHeader
}

func (c columnHeaderOption) getType() csvOptionsType {
	return columnHeader
}

type commaOption struct {
	comma rune
}

func CommaOption(comma rune) commaOption {
	return commaOption{
		comma: comma,
	}
}

func (c commaOption) getType() csvOptionsType {
	return comma
}

type commentOption struct {
	comment rune
}

func CommentOption(comment rune) commentOption {
	return commentOption{
		comment: comment,
	}
}

func (c commentOption) getType() csvOptionsType {
	return comment
}

type fieldsPerRecordOption struct {
	fieldsPerRecord int
}

func FieldPerRecordOption(i int) fieldsPerRecordOption {
	return fieldsPerRecordOption{
		fieldsPerRecord: i,
	}
}

func (f fieldsPerRecordOption) getType() csvOptionsType {
	return fieldsPerRecord
}

type lazyQuotesOption struct {
	lazyQuotes bool
}

func LazyQuotesOption(b bool) lazyQuotesOption {
	return lazyQuotesOption{
		lazyQuotes: b,
	}
}

func (l lazyQuotesOption) getType() csvOptionsType {
	return lazyQuotes
}

type trimLeadingSpaceOption struct {
	trimLeadingSpace bool
}

func TrimLeadingSpaceOption(b bool) trimLeadingSpaceOption {
	return trimLeadingSpaceOption{
		trimLeadingSpace: b,
	}
}

func (t trimLeadingSpaceOption) getType() csvOptionsType {
	return trimLeadingSpace
}

type reuseRecordOption struct {
	reuseRecord bool
}

func ReuseRecordOption(b bool) reuseRecordOption {
	return reuseRecordOption{
		reuseRecord: b,
	}
}

func (r reuseRecordOption) getType() csvOptionsType {
	return reuseRecord
}

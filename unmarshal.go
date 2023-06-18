package gocsvparser

import (
	"encoding/csv"
)

type Unmarshaler struct {
	header []string
	reader *csv.Reader
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

func (u *Unmarshaler) Unmarshal(data []byte, v interface{}) error {
	//TODO: implementation
	return nil
}

func Unmarshal(data []byte, v interface{}) error {
	return NewUnmarshaler().Unmarshal(data, v)
}

package gocsvparser

import (
	"reflect"
	"testing"
)

type Coba struct {
	Name    string `csv:"name"`
	Address string `csv:"address,omitempty"`
	Mile    int64  `csv:"mile"`
	anon
}

type anon struct {
	FieldX  int64 `csv:"x"`
	OutputY int64 `csv:"y"`
}

func TestParse(t *testing.T) {
	newDefaultHandler().detType(&Coba{})
}

func TestRead(t *testing.T) {
	var coba []Coba

	Read(&coba, t)

	t.Logf("outside %+v", coba)

	// type args struct {
	// 	i interface{}
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		Read(tt.args.i)
	// 	})
	// }
}

func Read(i interface{}, t *testing.T) {
	var val reflect.Value

	val = reflect.ValueOf(i)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()

		if val.Kind() == reflect.Slice {
			vslice := val
			typ := val.Type().Elem()

			if typ.Kind() == reflect.Struct {
				for _, x := range reflect.VisibleFields(typ) {
					t.Logf("test: %+v", x)
				}

				nv := reflect.New(typ).Elem()
				nv.FieldByName("Name").SetString("hello")
				nv.FieldByName("Mile").SetInt(72)
				// vslice = reflect.Append(vslice, nv)
				vslice.Set(reflect.Append(vslice, nv))

				t.Logf("inside: %+v", vslice)
			}
		}
	}
}

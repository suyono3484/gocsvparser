package gocsvparser_test

import (
	"testing"

	"github.com/suyono3484/gocsvparser"
)

type testDirect struct {
}

func (t *testDirect) SetFieldConfigs(configs []gocsvparser.FieldsConfig) {

}

func (t *testDirect) HandleRecord(v any, records []string) error {
	return nil
}

func TestExt(t *testing.T) {
	// _, err := gocsvparser.NewUnmarshaler().WithRecordHandler(new(testDirect))
	// if err != nil {
	// 	t.Fatalf("unexpected error %v", err)
	// }
}

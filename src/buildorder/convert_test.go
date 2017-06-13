package buildorder

import (
	"reflect"
	"testing"
)

func TestBuildOrder(t *testing.T) {
	buildOrder, err := Get("./testpackages")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(buildOrder, []string{
		"a",
		"b",
		"d",
		"c",
	}) {
		t.Fail()
	}
}

package blocks

import "testing"
import "reflect"

func TestNewGrassBlock(t *testing.T) {
	blocks := []interface{}{NewGrassBlock(), NewDirtBlock()}
	for _, v := range blocks {
		t.Log(reflect.TypeOf(v).String());
	}
}

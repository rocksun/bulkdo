package bulkdo

import (
	"strings"
	"testing"
)

func Test_readItems(t *testing.T) {
	itemStrs := strings.NewReader(`t1,t2,t3
a1,a2,a3
b1,b2,b3`)
	items, err := readItems(itemStrs)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected length is 2, but got %v", len(items))
	}

	expectedItem1T1 := "a1"
	if items[0]["t1"] != expectedItem1T1 {
		t.Errorf("expectedItem1T1 is %v, but got %v", expectedItem1T1, items[0]["t1"])
	}

	expectedItem2T2 := "b2"
	if items[1]["t2"] != expectedItem2T2 {
		t.Errorf("expectedItem2T2 is %v, but got %v", expectedItem2T2, items[1]["t2"])
	}
}

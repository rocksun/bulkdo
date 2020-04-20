package bulkdo

import (
	"errors"
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

type MockReader struct {
}

func (r *MockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Mock Error")
}

func Test_readItemsWithError(t *testing.T) {
	itemStrs := &MockReader{}
	_, err := readItems(itemStrs)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func Test_parseCommands(t *testing.T) {
	itemStrs := strings.NewReader(`t1,t2,t3
a1,a2,a3
b1,b2,b3`)
	items, _ := readItems(itemStrs)
	template := strings.NewReader("echo {{ .v.t1 }} {{ .v.t2 }}")
	commands, err := parseCommands(template, items)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if len(commands) != 2 {
		t.Errorf("Expected length is 2, but got %v", len(commands))
	}

	expected1 := "echo a1 a2"
	if expected1 != commands[0] {
		t.Errorf("Expected is '%v', but got '%v'", expected1, commands[0])
	}
}

func Test_execCommands(t *testing.T) {
	cmds := []string{`@echo off
echo a1
echo b1`,
		`@echo off
echo a2
echo b2`,
		`@echo off
echo a3
echo b3`,
	}

	outs, err := execCommands(cmds)
	if err != nil {
		t.Errorf("Exepected no err, but got %v", err)
	}

	if len(outs) != 3 {
		t.Errorf("Expected length is 3, but got %v", len(outs))
	}

	expected2 := "a2\r\nb2\r\n"
	if expected2 != outs[1] {
		t.Errorf("Expected '%v', but got '%v'", expected2, outs[1])
	}

}

package test

import (
	"testing"

	"github.com/mallsbill/device-detector-go/parser"
)

func TestParse(t *testing.T) {
	o := parser.NewOperatingSystem("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.81 Safari/537.36")
	res := o.Parse()

	if res != true {
		t.Error("parsing result not true", res)
	}
}

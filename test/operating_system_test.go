package test

import (
	"testing"

	"github.com/mallsbill/device-detector-go/parser"
)

func TestParseWindows(t *testing.T) {
	o := parser.NewOperatingSystem("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.81 Safari/537.36")
	res := o.Parse()

	if res != true {
		t.Error("parsing result not true", res)
	}

	if o.GetName() != "Windows" {
		t.Error("name is not equal to Windows", res)
	}

	if o.GetShort() != "WIN" {
		t.Error("short is not equal to WIN", res)
	}

	if o.GetVersion() != "10" {
		t.Error("version is not equal to 10", res)
	}

	if o.GetPlatform() != "x64" {
		t.Error("platform is not equal to x64", res)
	}

	if o.GetFamily() != "Windows" {
		t.Error("family is not equal to Windows", res)
	}
}

func TestParseCentos(t *testing.T) {
	o := parser.NewOperatingSystem("Mozilla/5.0 (X11; U; Linux i686 (x86_64); en-US; rv:1.9.0.6) Gecko/2009020414 CentOS/3.0.6-1.el5.centos Firefox/3.0.6")
	res := o.Parse()

	if res != true {
		t.Error("parsing result not true", res)
	}
}

func TestParseNothingToFound(t *testing.T) {
	o := parser.NewOperatingSystem("NothingToFound")
	res := o.Parse()

	if res != false {
		t.Error("parsing result not true", res)
	}
}

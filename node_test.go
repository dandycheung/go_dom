package dom

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

var b, ioErr = ioutil.ReadFile("./test/sample_html_1.html")

var doc, readerErr = html.Parse(bytes.NewReader(b))

// TestParseStyleNodeBody
func TestParseStyleNodeBody(t *testing.T) {
	n := QuerySelector("style", doc)

	if n == nil {
		t.Errorf("TestParseStyleNodeBody returned incorrect nil")
	}

	styles := ParseStyleNodeBody(n)

	if styles == nil {
		t.Errorf("TestParseStyleNodeBody returned incorrect nil")
	}

	// Init a new string map if selector name does not exist in map
	if _, ok := styles["btn"]; !ok {
		t.Errorf("TestParseStyleNodeBody returned incorrect nil")
	}

	// Init a new string map if selector name does not exist in map
	if _, ok := styles["divider"]; !ok {
		t.Errorf("TestParseStyleNodeBody returned incorrect nil")
	}

	// Init a new string map if selector name does not exist in map
	if _, ok := styles["header"]; !ok {
		t.Errorf("TestParseStyleNodeBody returned incorrect nil")
	}
}

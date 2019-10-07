package dom

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

var b, ioErr = ioutil.ReadFile("./test/sample_html_1.html")

var doc, readerErr = html.Parse(bytes.NewReader(b))

// TestParseStyleData
func TestParseStyleData(t *testing.T) {
	n := QuerySelector("style", doc)

	if n == nil {
		t.Errorf("TestParseStyleData returned incorrect nil")
	}

	styles := ParseStyleData(n)

	if styles == nil {
		t.Errorf("TestParseStyleData returned incorrect nil")
	}

	// Init a new string map if selector name does not exist in map
	if _, ok := styles["btn"]; !ok {
		t.Errorf("TestParseStyleData returned incorrect nil")
	}

	// Init a new string map if selector name does not exist in map
	if _, ok := styles["divider"]; !ok {
		t.Errorf("TestParseStyleData returned incorrect nil")
	}

	// Init a new string map if selector name does not exist in map
	if _, ok := styles["header"]; !ok {
		t.Errorf("TestParseStyleData returned incorrect nil")
	}
}

// // TestGetElementByClassnameMultiple func
// func TestGetElementByClassnameMultiple(t *testing.T) {
// 	el := GetElementByClassname("header-links primary-nav", doc)

// 	if el == nil {
// 		t.Errorf("GetElementByClassname returned incorrect nil")
// 	}
// }

// // TestGetElementByClassnameInvalid func
// func TestGetElementByClassnameInvalid(t *testing.T) {
// 	el := GetElementByClassname("header-links invalid-class", doc)

// 	if el != nil {
// 		t.Errorf("GetElementByClassname returned incorrect value.")
// 	}
// }

// // TestGetElementsByClassname func
// func TestGetElementsByClassname(t *testing.T) {
// 	n := GetElementsByClassname("section", doc)

// 	if n == nil {
// 		t.Errorf("GetElementsByClassname returned incorrect nil")
// 	}

// 	if len(n) != 2 {
// 		t.Errorf("GetElementsByClassname did not return 2 nodes")
// 	}
// }

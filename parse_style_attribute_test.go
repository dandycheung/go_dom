package dom

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

// TestParseStyleAttribute
func TestParseStyleAttribute(t *testing.T) {
	b, _ := ioutil.ReadFile("./test/sample_html_1.html")
	doc, _ := html.Parse(bytes.NewReader(b))
	body := GetDocumentBody(doc)

	styles := ParseStyleAttribute(body)

	if styles["display"] != "block" {
		t.Errorf("TestParseStyleAttribute \"display\" returned incorrect val")
	}

	if styles["background-color"] != "purple" {
		t.Errorf("TestParseStyleAttribute \"display\" returned incorrect val")
	}
}

package dom

// import (
// 	"bytes"
// 	"io/ioutil"
// 	"testing"

// 	"golang.org/x/net/html"
// )

// var b, ioErr = ioutil.ReadFile("./test_doc.html")

// var doc, readerErr = html.Parse(bytes.NewReader(b))

// // TestGetElementByClassname func
// func TestGetElementByClassname(t *testing.T) {
// 	n := GetElementByClassname("header-links", doc)

// 	if n == nil {
// 		t.Errorf("GetElementByClassname returned incorrect nil")
// 	}
// }

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

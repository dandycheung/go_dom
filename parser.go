package dom

import "golang.org/x/net/html"

// ParseImageSrc takes an html node and returns the src attribute
// of the first image element it finds.
func ParseImageSrc(n *html.Node) string {
	fn := func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "image" {
			return true
		}
		return false
	}

	img := FindOne(n, fn)

	if img == nil {
		return ""
	}

	return GetAttribute(img, "src")
}

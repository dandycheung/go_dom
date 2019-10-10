package dom

import (
	"strings"

	"golang.org/x/net/html"
)

// ParseStyleAttribute parses a node attribute styles into a map
// of key values.
func ParseStyleAttribute(n *html.Node) (styles map[string]string) {
	str := GetAttribute(n, "style")

	if str == "" {
		return styles
	}

	// init style map
	styles = make(map[string]string)

	// Split on semicolon separators
	split := strings.Split(str, ";")

	for _, body := range split {
		parts := strings.Split(strings.TrimSpace(body), ":")

		if len(parts) != 2 {
			continue
		}

		styles[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return styles
}

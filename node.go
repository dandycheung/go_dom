package dom

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// StyleDeclarations represents a map of classnames that maps to
// a map of style key values
type StyleDeclarations map[string]map[string]string

// GetAttribute accepts an *html.Node and a string key and returns the attribute value
func GetAttribute(n *html.Node, k string) string {
	var s string

	for _, a := range n.Attr {
		if a.Key == k {
			s = a.Val
			break
		}
	}

	return s
}

// ParseStyleNodeBody takes a style node and parses the data into a map of
// classnames and attribute key / values
func ParseStyleNodeBody(n *html.Node) StyleDeclarations {
	styleMap := make(StyleDeclarations)

	// require node to be a style node
	if n.Type != html.ElementNode || n.Data != "style" {
		return styleMap
	}

	// Create a new scanner to read each style line
	scanner := bufio.NewScanner(strings.NewReader(n.FirstChild.Data))
	lines := make([]string, 0)

	// read from scanner into slice
	for scanner.Scan() {
		txt := scanner.Text()
		txt = strings.TrimSpace(txt)

		if IsNotEmpty(txt) {
			lines = append(lines, txt)
		}
	}

	// If err, print and exit
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return styleMap
	}

	// Iterate each line separately e.g. ".btn{display:none;font-size:12px}"
	// and break the class selector and style definition into separate parts.
	for _, line := range lines {
		ch := line[:1]

		// We require the first character to start with a class selector,
		// support for additional selectors will need to be implemented later.
		if ch != "." {
			continue
		}

		// Get the starting index of the body declaration
		i := strings.Index(line, "{")

		// Exit if body declration is not found
		if i < 0 {
			continue
		}

		// parse the selector name
		sName := line[1:i]

		// Init a new string map if selector name does not exist in map
		if _, ok := styleMap[sName]; !ok {
			styleMap[sName] = make(map[string]string)
		}

		// parse body declaration from string
		body := line[i+1 : len(line)-1]

		// iterate over each rule in the body declaration and add it to the style map
		for _, rule := range strings.Split(body, ";") {
			trimRule := strings.TrimSpace(rule)

			if !IsNotEmpty(trimRule) {
				continue
			}

			// break rule into an array of two parts, key && val
			parts := make([]string, 0)

			// Split the rule on key value
			for _, val := range strings.Split(trimRule, ":") {
				trimVal := strings.TrimSpace(val)

				if !IsNotEmpty(trimVal) {
					continue
				}

				parts = append(parts, trimVal)
			}

			if len(parts) != 2 {
				continue
			}

			// Assign rule key value to style map
			styleMap[sName][parts[0]] = parts[1]
		}
	}

	return styleMap
}

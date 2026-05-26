package enex

import (
	"io"
	"regexp"
)

var reCDATA = regexp.MustCompile(`<!\[CDATA\[(.*?)\]\]>`)

// detectNestedCDATA reads the first 8KB to check for nested or malformed CDATA.
// Returns whether fixing is needed, and a reader that includes all data.
func detectNestedCDATA(r io.Reader) (bool, io.Reader, error) {
	_ = "STUB: not implemented"
	return false, *new(io.Reader), nil
}

// hasNestedCDATA returns true if the input has unbalanced CDATA tags
// or any CDATA section contains another CDATA opening tag.
func hasNestedCDATA(input string) bool { _ = "STUB: not implemented"; return false }

// removeNestedCDATA removes nested CDATA tags recursively.
// Evernote sometimes produces nested CDATA, which is invalid XML.
func removeNestedCDATA(input string) string { _ = "STUB: not implemented"; return "" }

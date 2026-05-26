package markdown

import (
	"github.com/mattn/godown"
)

// HighlightedText is a parsing rule to convert Evernote highlights to HTML spans with a background color
type HighlightedText struct{}

// Rule implements godown.CustomRule interface to extend basic conversion rules and
// convert text highlighted in Evernote to an inline HTML `span` tag with a custom background color
func (r *HighlightedText) Rule(next godown.WalkFunc) (string, godown.WalkFunc) {
	_ = "STUB: not implemented"
	return "", *new(godown.WalkFunc)
}

// TodoItem is a parsing rule to convert Evernote checkboxes to corresponding GitHub Flavoured Markdown items
type TodoItem struct{}

// Rule implements godown.CustomRule interface to handle Evernote-specific "en-todo" tag
// It converts the tag to a Markdown format with correct "checked" state
func (r TodoItem) Rule(next godown.WalkFunc) (string, godown.WalkFunc) {
	_ = "STUB: not implemented"
	return "", *new(godown.WalkFunc)
}

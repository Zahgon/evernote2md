package internal

import (
	"golang.org/x/net/html"

	"github.com/wormi4ok/evernote2md/encoding/enex"
	"github.com/wormi4ok/evernote2md/encoding/markdown"
)

// TagReplacer allows manipulating HTML nodes in order
// to present custom tags correctly in Markdown format after conversion
type TagReplacer interface {
	ReplaceTag(node *html.Node)
}

func (c *Converter) normalizeHTML(note *enex.Note, _ *markdown.Note, rr ...TagReplacer) {
	_ = "STUB: not implemented"
	return
}

// Media tag replacer puts a standard HTML <img> tag
// instead of custom <en-media> tag if it is an image
// and <a> tag for everything else to be able to download it as a file
type Media struct {
	resources map[string]markdown.Resource

	// If identifiers are missing we use resources one by one
	cnt int
}

var htmlFormat = map[markdown.ResourceType]string{
	markdown.Image: `<img src="%s/%s" alt="%s" />`,
	markdown.File:  `<a href="./%s/%s">%s</a>`,
}

// NewReplacerMedia creates a Media TagReplacer using resources as a data source
func NewReplacerMedia(resources map[string]markdown.Resource) *Media {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceTag implements the TagReplacer interface
func (r *Media) ReplaceTag(n *html.Node) { _ = "STUB: not implemented"; return }

func isMedia(n *html.Node) bool { _ = "STUB: not implemented"; return false }

func hashAttr(n *html.Node) string { _ = "STUB: not implemented"; return "" }

func replaceNode(n *html.Node, res markdown.Resource) { _ = "STUB: not implemented"; return }

func appendMedia(node, media *html.Node) { _ = "STUB: not implemented"; return }

// newline

// Since we control input, this wrapper gives a simple
// interface which will panic in case of bad strings
func parseOne(h string, context *html.Node) *html.Node { _ = "STUB: not implemented"; return nil }

func resourceReference(res markdown.Resource) string { _ = "STUB: not implemented"; return "" }

// Code replaces div tag stylized to look like code blocks with an actual <pre> tag
type Code struct{}

// ReplaceTag implements the TagReplacer interface
func (r *Code) ReplaceTag(n *html.Node) { _ = "STUB: not implemented"; return }

func isCode(n *html.Node) bool { _ = "STUB: not implemented"; return false }

// ExtraDiv removes extra line break in tables and lists
type ExtraDiv struct{}

// ReplaceTag implements the TagReplacer interface
func (*ExtraDiv) ReplaceTag(n *html.Node) { _ = "STUB: not implemented"; return }

func hasExtraDiv(n *html.Node) bool { _ = "STUB: not implemented"; return false }

// TextFormatter catches bold and italic, bold takes precedence
type TextFormatter struct{}

// ReplaceTag implements the TagReplacer interface
func (*TextFormatter) ReplaceTag(n *html.Node) { _ = "STUB: not implemented"; return }

func isBold(n *html.Node) bool { _ = "STUB: not implemented"; return false }

func isItalic(n *html.Node) bool { _ = "STUB: not implemented"; return false }

// EmptyAnchor removes anchor tags without text
type EmptyAnchor struct{}

// ReplaceTag implements the TagReplacer interface
func (*EmptyAnchor) ReplaceTag(n *html.Node) { _ = "STUB: not implemented"; return }

// NormalizeTodo replaces style-based checkboxes with tag-based checkboxes.
type NormalizeTodo struct{}

// ReplaceTag implements the TagReplacer interface
func (*NormalizeTodo) ReplaceTag(n *html.Node) { _ = "STUB: not implemented"; return }

// Figure out whether this item is checked or not.

// Construct a new <en-todo/> node, stealing the current node's children.

// Insert <en-todo/> node as only child of the parent node.

package internal

import (
	"regexp"

	"github.com/wormi4ok/evernote2md/encoding/enex"
	"github.com/wormi4ok/evernote2md/encoding/markdown"
)

// DefaultTagTemplate format if none specified
const DefaultTagTemplate = "`{{tag}}`"

const tagToken = "{{tag}}"

var spaces = regexp.MustCompile(`\s+`)

func (c *Converter) prependTags(note *enex.Note, md *markdown.Note) {
	_ = "STUB: not implemented"
	return
}

func (c *Converter) tagList(note *enex.Note, tagTemplate string, joinString string, spacesToUnderscores bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Default tag template allows spaces in tags, but for custom templates
// we replace all spaces with underscores to prevent word splitting

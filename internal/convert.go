package internal

import (
	"time"

	"github.com/wormi4ok/evernote2md/encoding/enex"
	"github.com/wormi4ok/evernote2md/encoding/markdown"
)

const FrontMatterTemplate = `---
date: '{{.CTime}}'
updated_at: '{{.MTime}}'
title: {{ trim .Title | quote }}
{{- if .TagList }}
tags: [ {{ .TagList }} ]
{{- end -}}
{{- with .Attributes -}}
{{- if .SourceUrl }}
url: {{ trim .SourceUrl -}}
{{- end -}}
{{- if .Latitude }}
latitude: {{ .Latitude -}}
{{- end -}}
{{- if .Longitude }}
longitude: {{ .Longitude -}}
{{- end -}}
{{- if .Altitude }}
altitude: {{ .Altitude -}}
{{- end -}}
{{- if .Source }}
source: {{ trim .Source -}}
{{- end }}
{{- end }}

---

`

// Converter holds configuration options to control conversion
type Converter struct {
	TagTemplate         string
	EnableHighlights    bool
	EscapeSpecialChars  bool
	EnableFrontMatter   bool
	FrontMatterTemplate string

	// err holds an error during conversion
	// Every conversion step should check this field and skip execution if it is not empty
	err error
}

// NewConverter creates a Converter with valid tagTemplate
func NewConverter(tagTemplate string, enableFrontMatter, enableHighlights, escapeSpecialChars bool) (*Converter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert an Evernote file to markdown
func (c *Converter) Convert(note *enex.Note) (*markdown.Note, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Converter) mapResources(note *enex.Note, md *markdown.Note) {
	_ = "STUB: not implemented"
	return
}

// Ensure the name is unique

func (c *Converter) prependTitle(note *enex.Note, md *markdown.Note) {
	_ = "STUB: not implemented"
	return
}

func (c *Converter) toMarkdown(note *enex.Note, md *markdown.Note) {
	_ = "STUB: not implemented"
	return
}

func (c *Converter) trimSpaces(_ *enex.Note, md *markdown.Note) { _ = "STUB: not implemented"; return }

func (c *Converter) addDates(note *enex.Note, md *markdown.Note) { _ = "STUB: not implemented"; return }

const dateFrontMatterFormat = "2006-01-02 15:04:05 -0700"

func (c *Converter) addFrontMatter(note *enex.Note, md *markdown.Note) {
	_ = "STUB: not implemented"
	return
}

const evernoteDateFormat = "20060102T150405Z"

// 20180109T173725Z -> 2018-01-09T17:37:25Z
func convertEvernoteDate(evernoteDate string) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

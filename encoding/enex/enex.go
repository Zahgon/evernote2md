package enex

import (
	"encoding/xml"
	"io"
	"regexp"
)

type (
	// Export represents Evernote enex file structure
	Export struct {
		XMLName xml.Name `xml:"en-export"`
		Date    string   `xml:"export-date,attr"`
		Notes   []Note   `xml:"note"`
	}

	// Note is one note in Evernote
	Note struct {
		XMLName    xml.Name       `xml:"note"`
		Title      string         `xml:"title"`
		Content    []byte         `xml:"content"`
		Updated    string         `xml:"updated"`
		Created    string         `xml:"created"`
		Tags       []string       `xml:"tag"`
		Attributes NoteAttributes `xml:"note-attributes"`
		Resources  []Resource     `xml:"resource"`
	}

	// NoteAttributes contain the note metadata
	NoteAttributes struct {
		Source            string `xml:"source"`
		SourceApplication string `xml:"source-application"`
		Latitude          string `xml:"latitude"`
		Longitude         string `xml:"longitude"`
		Altitude          string `xml:"altitude"`
		Author            string `xml:"author"`
		SourceUrl         string `xml:"source-url"`
	}

	// Resource embedded in the note
	Resource struct {
		ID          string
		Type        string
		Data        Data       `xml:"data"`
		Mime        string     `xml:"mime"`
		Width       int        `xml:"width"`
		Height      int        `xml:"height"`
		Attributes  Attributes `xml:"resource-attributes"`
		Recognition []byte     `xml:"recognition"`
	}

	// Attributes of the resource
	Attributes struct {
		Timestamp string `xml:"timestamp"`
		Filename  string `xml:"file-name"`
		SourceUrl string `xml:"source-url"`
	}
	// Recognition for the resource
	Recognition struct {
		XMLName xml.Name `xml:"recoIndex"`
		ObjID   string   `xml:"objID,attr"`
		ObjType string   `xml:"objType,attr"`
	}

	// Data object in base64
	Data struct {
		XMLName  xml.Name `xml:"data"`
		Encoding string   `xml:"encoding,attr"`
		Content  []byte   `xml:",innerxml"`
	}

	// Content of Evernote Notes
	Content struct {
		Text []byte `xml:",innerxml"`
	}
)

var hashRe = regexp.MustCompile(`\b[0-9a-f]{32}\b`)

// Decode will return an Export from evernote
func Decode(data io.Reader) (*Export, error) { _ = "STUB: not implemented"; return nil, nil }

// EOF is a known case when the content is empty

type Decoder struct {
	xml *xml.Decoder
}

func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

func (d Decoder) Decode(v any) error { _ = "STUB: not implemented"; return nil }

type StreamDecoder struct {
	xml *xml.Decoder
}

func NewStreamDecoder(r io.Reader) (*StreamDecoder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d StreamDecoder) Next(n *Note) error { _ = "STUB: not implemented"; return nil }

func decodeContent(n *Note) error { _ = "STUB: not implemented"; return nil }

func decodeRecognition(n *Note) error { _ = "STUB: not implemented"; return nil }

// findEnExportElement advances the decoder to the en-export element.
func findEnExportElement(decoder *xml.Decoder) error { _ = "STUB: not implemented"; return nil }

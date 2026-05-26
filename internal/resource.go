package internal

import (
	"io"
	"regexp"

	"github.com/wormi4ok/evernote2md/encoding/enex"
)

var reImg = regexp.MustCompile(`^image/[\w]+`)

var reFileAndExt = regexp.MustCompile(`(.*)(\.[\w\d]+)`)

var reBase64 = regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)

func decoder(d enex.Data) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func isBase64Encoded(content []byte) bool { _ = "STUB: not implemented"; return false }

func isImage(mimeType string) bool { _ = "STUB: not implemented"; return false }

func name(r enex.Resource) (name string, extension string) {
	_ = "STUB: not implemented"

	// Try to split a file into name and extension
	return "", ""
}

// Guess the extension by the mime type

// Return sanitized filename

// guessName of the res with the following priority:
// 1. Filename attribute
// 2. SourceUrl attribute (but use ID when sourceUrl starts with "en-cache://")
// 3. ID of the res (hash)
// 4. File type as name
func guessName(r enex.Resource) string { _ = "STUB: not implemented"; return "" }

var preferredExt = map[string]string{
	"image/jpeg": ".jpg",
}

func guessExt(mimeType string) string { _ = "STUB: not implemented"; return "" }

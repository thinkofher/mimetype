// Package mimetype uses magic number signatures to detect the MIME type of a file.
//
// mimetype stores the list of MIME types in a tree structure with
// "application/octet-stream" at the root of the hierarchy. The hierarchy
// approach minimizes the number of checks that need to be done on the input
// and allows for more precise results once the base type of file has been
// identified.
package mimetype

import (
	"io"
	"os"

	"github.com/gabriel-vasile/mimetype/internal/matchers"
)

// Detect returns the MIME type found from the provided byte slice.
//
// Failure to identify the format results in application/octet-stream
// being returned.
func Detect(in []byte) (mime *MIME) {
	if len(in) == 0 {
		return newMIME("inode/x-empty", "", matchers.True)
	}

	return root.match(in, root)
}

// DetectReader returns the MIME type of the provided reader.
//
// The result is always a valid MIME type, with application/octet-stream
// returned when identification failed with or without an error.
// Any error returned is related to the reading from the input reader.
func DetectReader(r io.Reader) (mime *MIME, err error) {
	in := make([]byte, matchers.ReadLimit)
	n, err := r.Read(in)
	if err != nil && err != io.EOF {
		return root, err
	}
	in = in[:n]

	return Detect(in), nil
}

// DetectFile returns the MIME type of the provided file.
//
// The result is always a valid MIME type, with application/octet-stream
// returned when identification failed with or without an error.
// Any error returned is related to the opening and reading from the input file.
func DetectFile(file string) (mime *MIME, err error) {
	f, err := os.Open(file)
	if err != nil {
		return root, err
	}
	defer f.Close()

	return DetectReader(f)
}

package mimetype_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

// To find the MIME type of some bytes/reader/file, perform a detect on the input.
func Example_detect() {
	file := "testdata/pdf.pdf"
	reader, _ := os.Open(file)
	data, _ := ioutil.ReadFile(file)

	dmime := mimetype.Detect(data)
	rmime, rerr := mimetype.DetectReader(reader)
	fmime, ferr := mimetype.DetectFile(file)

	fmt.Println(dmime, rmime, fmime)
	fmt.Println(rerr, ferr)

	// Output: application/pdf application/pdf application/pdf
	// <nil> <nil>
}

// To check if some bytes/reader/file has a specific MIME type, first perform
// a detect on the input and then test against the MIME. `Is` also works with
// MIME aliases.
func Example_check() {
	mime, err := mimetype.DetectFile("testdata/zip.zip")
	// application/x-zip is an alias of application/zip,
	// therefore Is returns true both times.
	fmt.Println(mime.Is("application/zip"), mime.Is("application/x-zip"), err)

	// Output: true true <nil>
}

// To check if some bytes/reader/file has a base MIME type, first perform
// a detect on the input and then navigate the parents until the base MIME type
// is found.
func Example_parent() {
	// Ex: if you are interested in text/plain and all of its subtypes:
	// text/html, text/xml, text/csv, etc.
	mime, err := mimetype.DetectFile("testdata/html.html")

	isText := false
	for ; mime != nil; mime = mime.Parent() {
		if mime.Is("text/plain") {
			isText = true
		}
	}

	// isText is true, even if the detected MIME was text/html.
	fmt.Println(isText, err)

	// Output: true <nil>
}

func ExampleDetect() {
	data, err := ioutil.ReadFile("testdata/zip.zip")
	mime := mimetype.Detect(data)

	fmt.Println(mime, err)

	// Output: application/zip <nil>
}

func ExampleDetectReader() {
	data, oerr := os.Open("testdata/zip.zip")
	mime, merr := mimetype.DetectReader(data)

	fmt.Println(mime, oerr, merr)

	// Output: application/zip <nil> <nil>
}

func ExampleDetectFile() {
	mime, err := mimetype.DetectFile("testdata/zip.zip")

	fmt.Println(mime, err)

	// Output: application/zip <nil>
}

func ExampleMIME_Is() {
	mime, err := mimetype.DetectFile("testdata/pdf.pdf")

	pdf := mime.Is("application/pdf")
	xpdf := mime.Is("application/x-pdf")
	txt := mime.Is("text/plain")
	fmt.Println(pdf, xpdf, txt, err)

	// Output: true true false <nil>
}

package contenttype_test

import (
	"testing"

	"github.com/edwardofclt/go-contenttype"
	"github.com/karlseguin/expect"
)

func TestFileExtension(t *testing.T) {
	contentType, err := contenttype.DetectFile("test.js")
	expect.IsNil(err)
	expect.Expect(t, contentType, "text/javascript")
}

func TestNonExistentFileExtension(t *testing.T) {
	_, err := contenttype.DetectFile("test.foobar")
	expect.Expect(t, err, "open test.foobar: no such file or directory")
}

func TestApacheMimeTypeFileExtension(t *testing.T) {
	contentType, err := contenttype.DetectFile("test.cdmid")
	expect.IsNil(err)
	expect.Expect(t, contentType, "application/cdmi-domain")
}

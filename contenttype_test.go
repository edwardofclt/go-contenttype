package contenttype_test

import (
	"testing"

	"github.com/edwardofclt/go-contenttype"
)

func TestFileExtension(t *testing.T) {
	t.Log(contenttype.DetectFile("test.js"))
}

func TestNonExistentFileExtension(t *testing.T) {
	t.Log(contenttype.DetectFile("test.foobar"))
}

func TestApacheMimeTypeFileExtension(t *testing.T) {
	t.Log(contenttype.DetectFile("test.cdmid"))
}

package contenttype

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var knownfiles = []string{
	"/etc/mime.types",
	"/etc/httpd/mime.types",
	"/etc/httpd/conf/mime.types",
	"/etc/apache/mime.types",
	"/etc/apache2/mime.types",
	"/usr/local/etc/httpd/conf/mime.types",
	"/usr/local/lib/netscape/mime.types",
	"/usr/local/etc/httpd/conf/mime.types",
	"/usr/local/etc/mime.types",
}

func DetectFile(file string) (string, error) {
	// try by file extension using known file extensions
	extension := filepath.Ext(file)
	if contentType, known := Types[extension]; known && contentType != "" {
		return contentType, nil
	}

	for _, mimeTypeFile := range knownfiles {
		fileType, err := compareWithMimeFile(extension, mimeTypeFile)
		if err != nil {
			return "", err
		}

		if fileType != "" {
			return fileType, nil
		}
	}

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fileContent, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	finalAttempt := http.DetectContentType(fileContent)
	return finalAttempt, nil
}

func compareWithMimeFile(extension, mimeTypeFile string) (string, error) {
	f, err := os.Open(mimeTypeFile)
	if err != nil {
		// we don't throw an error because we can't assume the file will be there
		return "", nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// if the first character of the line is a comment, skip it
		if line[0] == '#' {
			continue
		}

		parts := strings.Split(scanner.Text(), "\t")
		if fmt.Sprintf(".%s", parts[len(parts)-1]) == extension {
			return parts[0], nil
		}
	}
	return "", nil
}

// Package untar untars a tarball to disk.
package untar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

// SpecificFile extracts the file with the specified filename from
// the tar.gz file in src and writes it to the file.
func SpecificFile(src io.Reader, filename string) {
	gz, err := gzip.NewReader(src)
	if err != nil {
		log.Fatal(err)
	}
	defer gz.Close()

	t := tar.NewReader(gz)
	for {
		header, err := t.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		// Extracts only regular file with the specified filename
		if !(header.Typeflag == tar.TypeReg && strings.Contains(header.Name, filename)) {
			continue
		}

		out, err := os.OpenFile(header.Name, os.O_WRONLY|os.O_CREATE, header.FileInfo().Mode())
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		if _, err = io.Copy(out, t); err != nil {
			log.Fatal(err)
		}
	}
}

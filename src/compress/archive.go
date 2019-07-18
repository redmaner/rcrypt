package compress

import (
	"archive/zip"
	"bytes"
)

type Archive struct {
	ZipWriter *zip.Writer
	Out       *bytes.Buffer
	isClosed  bool
}

func NewArchive() *Archive {

	// Create the zip archive
	newZipFile := new(bytes.Buffer)
	zipWriter := zip.NewWriter(newZipFile)

	return &Archive{
		ZipWriter: zipWriter,
		Out:       newZipFile,
	}
}

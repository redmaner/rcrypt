package compress

import (
	"archive/zip"
	"bytes"
)

type Archive struct {
	zipWriter *zip.Writer
	zipReader *zip.Reader
	out       *bytes.Buffer
	isClosed  bool
}

func NewArchive() *Archive {

	// Create the zip archive
	newZipFile := new(bytes.Buffer)
	zipWriter := zip.NewWriter(newZipFile)

	return &Archive{
		zipWriter: zipWriter,
		out:       newZipFile,
	}
}

func LoadArchive(data []byte) (*Archive, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}

	return &Archive{
		zipReader: zipReader,
		isClosed:  true,
	}, nil
}

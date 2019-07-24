package compress

import (
	"archive/zip"
	"bytes"
)

// Archive holds a zip archive that can either be created as new, or loaded
// from an existing archive
type Archive struct {
	zipWriter *zip.Writer
	zipReader *zip.Reader
	out       *bytes.Buffer
	isClosed  bool
}

// NewArchive returns a new created archive
func NewArchive() *Archive {

	// Create the zip archive
	newZipFile := new(bytes.Buffer)
	zipWriter := zip.NewWriter(newZipFile)

	return &Archive{
		zipWriter: zipWriter,
		out:       newZipFile,
	}
}

// LoadArchive loads an existing archive from a slice of bytes
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

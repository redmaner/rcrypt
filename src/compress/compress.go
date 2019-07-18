package compress

import (
	"fmt"
	"io"
	"os"
)

// close is used to close the archive. This should be done before compressing.
func (a *Archive) close() error {
	if a.isClosed {
		return fmt.Errorf("Archive is already compressed")
	}
	err := a.ZipWriter.Close()
	if err != nil {
		return fmt.Errorf("Error compressing archive: %v", err)
	}
	a.isClosed = true
	return nil
}

// Compress is used to compress the archive. It returns the zip in a slice of bytes.
func (a *Archive) Compress() ([]byte, error) {

	// Close the archive
	err := a.close()
	if err != nil {
		return []byte{}, err
	}

	// Return zip archive in bytes
	return a.Out.Bytes(), nil
}

// CompressToFile is used to compress the archive a to p file.
func (a *Archive) CompressToFile(p string) error {

	// Close the archive
	err := a.close()
	if err != nil {
		return err
	}

	// Create new zipfile
	zipFile, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("Creating zip %s failed: %v", p, err)
	}
	defer zipFile.Close()

	// Copy archive to file
	_, err = io.Copy(zipFile, a.Out)
	if err != nil {
		return fmt.Errorf("Error writing to zip %s: %v", p, err)
	}

	return nil
}

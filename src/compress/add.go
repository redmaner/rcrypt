package compress

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func (a *Archive) Add(p string) error {

	files := make([]string, 0, 8)

	// Create a slice of files that require to be zipped
	switch dir, err := os.Stat(p); {
	case err != nil:
		return err
	case dir.IsDir():
		files, err = walkDir(p)
		if err != nil {
			return err
		}
	default:
		files = append(files, p)
	}

	for _, f := range files {
		fileToZip, err := os.Open(f)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		// Get the file information
		info, err := fileToZip.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Using FileInfoHeader() above only uses the basename of the file. If we want
		// to preserve the folder structure we can overwrite this with the full path.
		header.Name = f

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := a.zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}
	}
	return nil
}

func walkDir(p string) ([]string, error) {
	files := make([]string, 0, 8)
	if err := filepath.Walk(p, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return files, err
	}
	return files, nil
}

package compress

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Decompress is a function to decompress the archive to dst, where dst is the
// destination directory
func (a *Archive) Decompress(dst string) error {

	for _, f := range a.zipReader.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dst, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			// Make Folder
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// Make File
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}

		// Close the file without defer to close before next iteration of loop
		if err := outFile.Close(); err != nil {
			return err
		}

		if err := rc.Close(); err != nil {
			return err
		}
	}

	return nil

}
